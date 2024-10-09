import requests
import uuid
import random

order_ids = []
for _ in range(1000):
    order_ids.append(uuid.uuid4().__str__())

customer_ids = []
for _ in range(100):
    customer_ids.append(uuid.uuid4().__str__())

for index in range(120):
    customer_id = random.choice(customer_ids)
    quantity = random.randint(1, 10)

    order_items = []
    for _ in range(quantity):
        order_items.append(
            {
                "id": random.choice(order_ids),
                "quantity": random.randint(1, 10),
                "price": random.randint(1, 10000),
            }
        )

    order = {
        "customerId": customer_id,
        "orderItems": order_items,
    }

    r = requests.post("http://localhost:3000/orders", json=order)
    r.status_code
    print(f"order {index + 1} created")