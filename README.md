fg-piface
=========

At last instead of dabbling yourself, u can share some of teh stuff ere with this code..

This is a bit of fun and some research into creating a MCP interface using
- a raspberry pi model B (NOT b+)
- a piface digital io board (shich does not fit B= without extender)
- a network cable
- Fg comms via websocket
- and golang for its channels
- or download a binary

- http://www.piface.org.uk/products/piface_digital/

and its working..
Threre is no a cunning plan to use golang..
and create a virtual "arinc" can bus..
- The golang as the "router"
  - weksocket interface
  - udp custom protocol interface
  - all transmit and udpate state on the varinc
- attaches nodes
  - these could be an arduino and its digit in, out
  - rpi board with a piface.. a display unit, and digital in/outs
  - Basically making a sim pit with switches and lights etc..
  
So the main problem is how to "communicate" with Fg
- and make all the lights go on and off etc,
- and presing a button makes things happen

So after some research..
with golang, and multiple channel go routines
its kinda working and being a boring olde progammer
- on the rpi thing.. the switches and ligts go on an off with display
  - and piface, digital io and inputs
- arduino -
  - is working as analog reads and major decicions
  
  Next step is to make things talk to each other and multiplex
  and for that we need i2c interface, ie single wire..
  
  


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

# Update 2015-06-27
  - multiplexed 16 indicators to leds using arduino, 3 pins and 2*595 shift register
  - rotary encoder to set var
  - 2 * seven segment with 13 pins !! (hard way)

# Update 2015-06-22
- The websocket and buttons work
- The config is now the problem in sending message




