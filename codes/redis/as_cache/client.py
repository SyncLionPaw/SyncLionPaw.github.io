import random
import time
from faker import Faker
from mockdb import DBClient

def send_random_requests(db_client: DBClient, num_requests: int):
    """Sends a specified number of random requests to the DBClient."""
    fake = Faker()
    for _ in range(num_requests):
        key = fake.word()  # Generate a random word for the key
        
        # Randomly choose between get and set operations
        if random.random() < 0.8:  # 50% chance of get
            value = db_client.get(key)
            print(f"GET key={key}, value={value}")
        else:  # 50% chance of set
            value = fake.word()  # Generate a random word for the value
            db_client.set(key, value)
            print(f"SET key={key}, value={value}")
        
        time.sleep(random.uniform(0.1, 0.5))  # Simulate some delay between requests

if __name__ == '__main__':
    db_client = DBClient()
    num_requests = 20
    send_random_requests(db_client, num_requests)