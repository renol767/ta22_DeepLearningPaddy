import random
import numpy as np

numbers1 = np.random.randint(0, 256, 3, dtype='uint8')
numbers = [[12, 133, 44], [184, 15, 15]]
color = None
color = color or [random.randint(0, 255) for _ in range(3)]
print(numbers1)
print(numbers)

print(color)