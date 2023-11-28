import json
import uuid
import random

node_count = 300
device_model = "tesla_t4"
vendor = "nvidia"
device_count_per_node = 8

devices = []

for i in range(node_count):
    for j in range(device_count_per_node):
        device = {
            "IP": f"10.0.{i // 256}.{i % 256}",
            "Index": j,
            "ID": str(uuid.uuid4()),
            "Model": device_model,
            "Vendor": "nvidia",
            "Memory": random.choice([2000, 4000, 6000, 7000, 8000, 11000])
        }
        devices.append(device)

with open(f"gpu_node_{node_count}.json", "w") as file:
    json.dump(devices, file, indent=2)
