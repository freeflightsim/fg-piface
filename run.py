#!/usr/bin/env python

import socket
import pifacedigitalio 
import time
import thread
import sys
import signal

exitapp = False

pif = pifacedigitalio.PiFaceDigital(init_board=True)


port  = 45670



class LED:
	ap = 0 # auto pilot light
	at = 1 # auto throttle light
	hdg_hold = 2 # heading hold light
	alt_hold = 3 # altitude hold light


def x_cb(ev):
	print("YES", ev.pin_num)
	pif.output_pins[1].value = not pif.output_pins[1].value

listener = pifacedigitalio.InputEventListener(chip=pif)
for i in range(0, 4):
	listener.register(i, pifacedigitalio.IODIR_FALLING_EDGE, x_cb)



def run_client():

	global port
	sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
	sock.bind( ("",
 port) )

	print "listenting on", port
	c = 0
	while True:
		data, addr = sock.recvfrom(1024)
		
		if c % 100 == 0:
			print "=", addr, data
		parts = data.replace("#", "").split("|")
		
		flags = [int(p) for p in parts]
	


		pif.output_pins[LED.ap].value = 1 if flags[0]  == 1 else 0
		pif.output_pins[LED.at].value = 1 if flags[1]  == 1 else 0
		pif.output_pins[LED.hdg_hold].value = 1 if flags[2]  == 1 else 0



#t = threading.Thread(target=udp_listen)


def start():
	threads = []
	t = UDPClient()
	t.start()
	threads.append(t)

	for tr in threads:
		tr.join()

def cleanup():
	print "cleanup"
	#sys.exit(0)




listener.activate()

run_client()
"""
try:
	thread.start_new_thread(run_client, ())
	#run_client()
except Exception, e:
	print e
"""
"""
if __name__ == "__main__":
	signal.signal(signal.SIGINT, cleanup)
	signal.signal(signal.SIGTERM, cleanup)
	start()	
"""	