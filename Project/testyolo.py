import cv2

con = 0.6464958190917969
cs = 'keyboard'
x1 = 127.39630889892578
y1 = 537.3564453125
x2 = 762.15185546875
y2 = 716.779052734375
img = cv2.imread('sd/ori/1647664924.7743068.jpg')
img = cv2.rectangle(img, (x1, y1), (x2, y2), (0, 255, 0), 2)
cv2.imshow('sda', img)
cv2.waitKey(1)