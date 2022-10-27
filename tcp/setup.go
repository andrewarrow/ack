package tcp

// TCP provides reliable, ordered, and error-checked delivery
// of a stream of octets (bytes)

/*
TCP is connection-oriented, and a connection between client and server
is established before data can be sent.

The server must be listening (passive open) for connection requests
from clients before a connection is established.

Three-way handshake (active open), retransmission, and error detection
adds to reliability but lengthens latency.

To avoid congestive collapse, TCP uses multi-faceted congestion-control strategy.

For each connection, TCP maintains a CWND, limiting the total number
of unacknowledged packets that may be in transit end-to-end.

This is somewhat analogous to TCP's sliding window used for flow control.

An application does not need to know the particular mechanisms
for sending data via a link to another host, such as the required IP
fragmentation to accommodate the maximum transmission unit of the
transmission medium.

At the transport layer, TCP handles all handshaking and transmission details
and presents an abstraction of the network connection to the application
typically through a network socket interface.


At the lower levels of the protocol stack, due to network congestion,
traffic load balancing, or unpredictable network behaviour,
IP packets may be lost, duplicated, or delivered out of order.

TCP detects these problems, requests re-transmission of lost data,
rearranges out-of-order data and even helps minimize network congestion
to reduce the occurrence of the other problems.

https://en.wikipedia.org/wiki/Transmission_Control_Protocol
*/
