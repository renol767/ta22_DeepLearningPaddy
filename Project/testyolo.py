import torch

# Model
model = torch.hub.load('ultralytics/yolov5', 'yolov5x')  # or yolov5m, yolov5l, yolov5x, custom

# Images
img = 'test.jpg'  # or file, Path, PIL, OpenCV, numpy, list

# Inference
results = model(img)

# Results
results.save(save_dir='results/')