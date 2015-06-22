fg-piface
=========

FlightGear leds and buttons with PiFace (soon 7 seg and knobs)

This is a bit of fun and some research into creating a MCP interface using
- a raspberry pi model B (NOT b+)
- a piface digital io board
- Fg comms via websocket
- and golang for its channels

http://www.piface.org.uk/products/piface_digital/

Run
==============

fgfs Host Machine
--------------------
- determine ip address of the host machine
  ```ifconfig```
- Start fgms on the full On machine with
  ```fgms --httpd=56789```




So Far
=============================

# Update 2015-06-22
- The websocket and buttons work
- The config is now the problem in sending message




