import json
import time

import matplotlib.pyplot as plt

def plot_counts_live(file_path="monitor.json", interval=1):
    """
    Continuously reads rcount and wcount from a JSON file and plots their change over time.

    Args:
        file_path (str): Path to the JSON file. Defaults to "monitor.json".
        interval (int): Time interval (in seconds) between updates. Defaults to 1.
    """
    rcounts = []
    wcounts = []
    time_points = []

    plt.ion()  # Turn on interactive mode for live plotting
    fig, ax = plt.subplots(figsize=(10, 6))
    line_r, = ax.plot([], [], label='Read Count (rcount)')
    line_w, = ax.plot([], [], label='Write Count (wcount)')

    ax.set_xlabel('Time')
    ax.set_ylabel('Count')
    ax.set_title('Read and Write Counts Over Time')
    ax.grid(True)
    ax.legend()

    start_time = time.time()  # Record the starting time

    try:
        while True:
            try:
                with open(file_path, 'r') as f:
                    data = json.load(f)
                rcount = data.get('rcount', 0)
                wcount = data.get('wcount', 0)

                if rcount is None or wcount is None:
                    print("Error: rcount or wcount data not found in JSON file.")
                    time.sleep(interval)
                    continue

                # Append the new data to the lists
                rcounts.append(rcount)
                wcounts.append(wcount)
                time_points.append(time.time() - start_time)

                # Update data on the plot
                line_r.set_xdata(time_points)
                line_r.set_ydata(rcounts)
                line_w.set_xdata(time_points)
                line_w.set_ydata(wcounts)

                # Adjust the plot limits to show the new data
                ax.set_xlim(min(time_points), max(time_points))
                ax.set_ylim(min(min(rcounts), min(wcounts)), max(max(rcounts), max(wcounts)))

                # Redraw the plot
                fig.canvas.draw()
                fig.canvas.flush_events()

            except FileNotFoundError:
                print(f"Error: File not found at {file_path}")
            except json.JSONDecodeError:
                print(f"Error: Invalid JSON format in {file_path}")
            except Exception as e:
                print(f"An unexpected error occurred: {e}")

            time.sleep(interval)

    except KeyboardInterrupt:
        print("Live plotting stopped.")

if __name__ == '__main__':
    plot_counts_live()