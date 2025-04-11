import socket
import threading
import argparse
import json


class DB:
    def __init__(self, db_file="db.txt"):
        self.db_file = db_file

    def set(self, key, value):
        with open(self.db_file, "a") as f:
            f.write(f"{key} {value}\n")
        return "OK"

    def get(self, key):
        with open(self.db_file, "r") as f:
            for line in f:
                k, v = line.strip().split(maxsplit=1)
                if k == key:
                    return v
        return None


class DBServer:
    def __init__(self, host="0.0.0.0", port=12345, db_file="db.txt", monitor_file="monitor.json"):
        self.host = host
        self.port = port
        self.db = DB(db_file)
        self.monitor_file = monitor_file
        self.rcount = 0
        self.wcount = 0
        self.load_counts()

    def load_counts(self):
        try:
            with open(self.monitor_file, "r") as f:
                data = json.load(f)
                self.rcount = data.get("rcount", 0)
                self.wcount = data.get("wcount", 0)
        except FileNotFoundError:
            self.rcount = 0
            self.wcount = 0
        except json.JSONDecodeError:
            self.rcount = 0
            self.wcount = 0

    def save_counts(self):
        data = {"rcount": self.rcount, "wcount": self.wcount}
        with open(self.monitor_file, "w") as f:
            json.dump(data, f)

    def handle_client(self, client_socket):
        with client_socket:
            while True:
                data = client_socket.recv(1024).decode("utf-8").strip()
                if not data:
                    break
                command, *args = data.split()
                if command == "SET" and len(args) == 2:
                    key, value = args
                    response = self.db.set(key, value)
                    self.wcount += 1
                    self.save_counts()
                    client_socket.sendall(f"{response}\n".encode("utf-8"))
                elif command == "GET" and len(args) == 1:
                    key = args[0]
                    value = self.db.get(key)
                    self.rcount += 1
                    self.save_counts()
                    if value is not None:
                        client_socket.sendall(f"{value}\n".encode("utf-8"))
                    else:
                        client_socket.sendall(b"NOT FOUND\n")
                else:
                    client_socket.sendall(b"ERROR\n")

    def start(self):
        server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server.bind((self.host, self.port))
        server.listen(5)
        print(f"Server listening on port {self.port}...")

        while True:
            client_socket, addr = server.accept()
            print(f"Connection from {addr}")
            client_handler = threading.Thread(
                target=self.handle_client, args=(client_socket,)
            )
            client_handler.start()



class DBClient:
    def __init__(self, host="0.0.0.0", port=12345):
        self.host = host
        self.port = port

    def send_command(self, command):
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
            client_socket.connect((self.host, self.port))
            client_socket.sendall(command.encode("utf-8"))
            response = client_socket.recv(1024).decode("utf-8").strip()
            return response

    def set(self, key, value):
        command = f"SET {key} {value}"
        return self.send_command(command)

    def get(self, key):
        command = f"GET {key}"
        return self.send_command(command)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Run DBServer or DBClient")
    parser.add_argument(
        "mode", choices=["server", "client"], help="Specify 'server' or 'client'"
    )
    parser.add_argument(
        "--key", help="Key for client command", required=False
    )  # Make key optional
    parser.add_argument("--value", help="Value for client SET command", required=False)
    parser.add_argument("--host", default="0.0.0.0", help="Host address")
    parser.add_argument("--port", type=int, default=12345, help="Port number")

    args = parser.parse_args()

    if args.mode == "server":
        server = DBServer(host=args.host, port=args.port)
        server.start()
    elif args.mode == "client":
        client = DBClient(host=args.host, port=args.port)
        if args.key and args.value:
            response = client.set(args.key, args.value)
            print(f"SET Response: {response}")
        elif args.key:
            response = client.get(args.key)
            print(f"GET Response: {response}")
        else:
            print("Please provide a key for GET or a key and value for SET.")

# python mockdb.py --key city --value xiamen client
# python mockdb.py --key city  client
