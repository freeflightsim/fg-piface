#!/usr/bin/env python
import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.bind( ("192.168.5.12", 4567) )

while True:
	data, addr = sock.recvfrom(1024)
	print "=", addr, data