# Networking
## OSI Model

| Layer No.     | Layer Name         | Description / Functions                                       | Example Protocols / Devices             |
|---------------|--------------------|----------------------------------------------------------------|------------------------------------------|
| 7 (Top)       | Application Layer   | - Provides services to user apps                               | HTTP, FTP, DNS, SMTP                     |
|               |                    | - Interface for user-facing software                           |                                          |
| 6             | Presentation Layer  | - Data translation, encryption, compression                    | SSL/TLS, JPEG, MP4, ASCII, JSON, XML     |
|               |                    | - Ensures readable data format                                 |                                          |
| 5             | Session Layer       | - Starts, manages, terminates sessions                         | NetBIOS, RPC, PPTP                       |
|               |                    | - Controls dialog between apps                                 |                                          |
| 4             | Transport Layer     | - Reliable delivery, sequencing, port addressing               | TCP, UDP                                 |
|               |                    | - Flow and error control                                       |                                          |
| 3             | Network Layer       | - Logical addressing, routing                                  | IP, ICMP, OSPF, BGP                      |
|               |                    | - Handles packet delivery across networks                      |                                          |
| 2             | Data Link Layer     | - MAC addressing, frame formatting, error detection            | Ethernet, ARP, PPP                       |
|               |                    | - Node-to-node reliable transfer                               |                                          |
| 1 (Bottom)    | Physical Layer      | - Transmission of raw bits via electrical/optical signals      | Cables, Hubs, Wi-Fi, NICs                |
|               |                    | - Physical medium used for communication                       |                                          |

## OPERATIONS IN EACH LAYER
| OSI Layer     | Sender Side Operations                                  | Receiver Side Operations                               |
|---------------|---------------------------------------------------------|--------------------------------------------------------|
| 7. Application| Web/email request, DNS lookup, file upload              | Render HTML, show email, save file                    |
| 6. Presentation| Encrypt, compress, convert to ASCII/JSON               | Decrypt, decompress, translate to readable format     |
| 5. Session    | Establish session, set dialog control                   | Accept session, maintain state, end session           |
| 4. Transport  | Segment data, assign port, add seq#, TCP/UDP headers    | Reorder segments, remove duplicates, verify data      |
| 3. Network    | Add source/destination IP, decide routing path          | Check IP, deliver to next hop or final destination    |
| 2. Data Link  | Create frame, add MAC address, error detection          | Check MAC, validate frame, remove header              |
| 1. Physical   | Convert bits to electrical/optical signal               | Convert signal back to bits, pass to Data Link layer  |


## DATA UNIT

| OSI Layer    | Data Unit Name |
| ------------ | -------------- |
| Application  | Data           |
| Presentation | Data           |
| Session      | Data           |
| Transport    | Segment        |
| Network      | Packet         |
| Data Link    | Frame          |
| Physical     | Bits           |

## Devices Working at Each Layer

| Layer        | Devices Involved            |
| ------------ | --------------------------- |
| Application  | Web browsers, email clients |
| Presentation | Encryption tools (SSL/TLS)  |
| Session      | Session managers, APIs      |
| Transport    | Firewalls, load balancers   |
| Network      | Routers                     |
| Data Link    | Switches, NICs              |
| Physical     | Cables, Hubs, Repeaters     |

## TCP/IP MODEL 
| OSI Model Layer         | ↔ | TCP/IP Model Layer         |
|-------------------------|----|----------------------------|
| Application Layer       | ↔  | Application Layer          |
| Presentation Layer      | ↔  | Application Layer          |
| Session Layer           | ↔  | Application Layer          |
| Transport Layer         | ↔  | Transport Layer            |
| Network Layer           | ↔  | Internet Layer             |
| Data Link Layer         | ↔  | Network Access Layer       |
| Physical Layer          | ↔  | Network Access Layer       |

Before diving deep into the TCP (Transmission Control Protocol), it's important to build a solid foundation in several networking and computer science concepts. 

Here's a structured list of prerequisites, organized from fundamental to slightly advanced topics:

1. Basic Computer Networking Concepts

What is a computer network?
Types of networks (LAN, WAN, MAN, PAN)
Network topologies (star, mesh, bus)
Client-Server vs Peer-to-Peer model
IP addressing basics (IPv4 and optionally IPv6)

2. OSI and TCP/IP Models

OSI 7-Layer Model (focus on Layer 3 and Layer 4)

TCP/IP 4-Layer Model

Understand how TCP fits into these models (Layer 4: Transport)


3. IP Layer Fundamentals (Layer 3)

IP Addressing and Subnetting
Packet structure of IP
Routing basics (how packets move across networks)
ICMP basics (ping, traceroute)


4. Ports and Protocols
What are ports?
Difference between TCP and UDP
Common port numbers (HTTP - 80, HTTPS - 443, etc.)
Well-known protocols (DNS, HTTP, FTP, etc.)

5. UDP (User Datagram Protocol)
Understand UDP as a connectionless protocol
Compare UDP vs TCP
Why TCP is reliable and UDP is not

6. Data Encapsulation and Decapsulation

Understand how data is wrapped and unwrapped as it moves through layers

See how TCP headers are added during encapsulation

7. Binary and Hexadecimal Understanding

IP headers, TCP headers are best understood when you can read:
Binary bit fields
Hex dumps in packet analyzers like Wireshark

8. Network Tools Basics

Tools like:
ping
traceroute
netstat
telnet or nc (netcat)
Wireshark for packet analysis

9. Basics of Operating Systems
How OS handles sockets and connections
System calls (like socket(), connect(), bind(), etc.)

10. Sockets Programming Basics (Optional but Recommended)

If you're coding:
Understand socket APIs (especially in C, Python, or Go)
Basics of creating TCP servers and clients

Once You Have These, You Can Study TCP In Depth:
TCP header structure and flags (SYN, ACK, FIN, etc.)
Three-way handshake and teardown
Flow control (window size, sliding window)
Congestion control (slow start, congestion avoidance)
Sequence and acknowledgment numbers
Retransmission mechanisms
TCP timers
TCP state machine (LISTEN, SYN_SENT, ESTABLISHED, etc.)
TCP performance tuning

