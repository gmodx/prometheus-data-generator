import json
import uuid
import random
from datetime import datetime, timedelta

node_count = 300
device_model = "tesla_t4"
vendor = "nvidia"
device_count_per_node = 8

end_time = datetime(2023, 11, 1)
days = 1
start_time = end_time - timedelta(days=days)
interval = timedelta(seconds=15)

samples = []

dt = start_time
while dt <= end_time:
    for i in range(node_count):
        for j in range(device_count_per_node):
            sample = {
                "IP": f"10.0.{i // 256}.{i % 256}",
                "Index": j,
                "ID": str(uuid.uuid4()),
                "Model": device_model,
                "Vendor": "nvidia",
                "Memory": random.choice([2000, 4000, 6000, 7000, 8000, 11000]),
                "Unix": int(dt.timestamp())
            }
            samples.append(sample)
    dt += interval

with open(f"samples_gpu_node_{node_count}.json", "w") as file:
    json.dump(samples, file, indent=2)
