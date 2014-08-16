#!/usr/bin/env python
import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.bind( ("0.0.0.0", 45670) )

c = 0
while True:
	c += 1
	data, addr = sock.recvfrom(1024)
	print "=",  c, data