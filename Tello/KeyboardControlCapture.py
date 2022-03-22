from djitellopy import tello
import KeyPressModule as kp
import time
import cv2


kp.init()
tel = tello.Tello()
tel.connect()
print(tel.get_battery())
global img
tel.streamon()

def getKeyboardInput():
    lr, fb, ud, yv = 0, 0, 0, 0
    speed = 50

    if kp.getKey("LEFT"): lr = -speed
    elif kp.getKey("RIGHT"): lr = speed

    if kp.getKey("UP"): fb = speed
    elif kp.getKey("DOWN"): fb = -speed

    if kp.getKey("w"): ud = speed
    elif kp.getKey("s"): ud = -speed

    if kp.getKey("a"): yv = -speed
    elif kp.getKey("d"): yv = speed

    if kp.getKey("q"): tel.land(); time.sleep(3)
    if kp.getKey("e"): tel.takeoff()

    if kp.getKey("z"):
        cv2.imwrite(f'Images/{time.time()}.jpg', img)
        time.sleep(0.3)

    return [lr, fb, ud, yv]

while True:
    vals = getKeyboardInput()
    tel.send_rc_control(vals[0], vals[1], vals[2], vals[3])
    
    img = tel.get_frame_read().frame
    cv2.imshow("Image", img)
    cv2.waitKey(1)