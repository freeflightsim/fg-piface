

// Input + Output to a running flightgear instance
package fgio


/*
The idea is to contain the FlightGear sim interface
in this package and speak with either
- websocket for simple listeners
- protocol udp for faster stuff

so far--
- Sends websocket messages ie Commands
- recieves websocket frames, and channel to process
 */


