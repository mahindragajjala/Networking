# Networking Guide

## OSI Model

| Layer No. | Layer Name         | Description / Functions                                       | Example Protocols / Devices             |
|-----------|--------------------|----------------------------------------------------------------|------------------------------------------|
| 7 (Top)   | Application Layer   | - Provides services to user apps<br>- Interface for user-facing software | HTTP, FTP, DNS, SMTP              |
| 6         | Presentation Layer  | - Data translation, encryption, compression<br>- Ensures readable data format | SSL/TLS, JPEG, MP4, ASCII, JSON, XML |
| 5         | Session Layer       | - Starts, manages, terminates sessions<br>- Controls dialog between apps | NetBIOS, RPC, PPTP                |
| 4         | Transport Layer     | - Reliable delivery, sequencing, port addressing<br>- Flow and error control | TCP, UDP                         |
| 3         | Network Layer       | - Logical addressing, routing<br>- Handles packet delivery across networks | IP, ICMP, OSPF, BGP              |
| 2         | Data Link Layer     | - MAC addressing, frame formatting, error detection<br>- Node-to-node reliable transfer | Ethernet, ARP, PPP              |
| 1 (Bottom)| Physical Layer      | - Transmission of raw bits via electrical/optical signals<br>- Physical medium used for communication | Cables, Hubs, Wi-Fi, NICs |

---

## Operations in Each OSI Layer

| OSI Layer     | Sender Side Operations                                  | Receiver Side Operations                               |
|---------------|---------------------------------------------------------|--------------------------------------------------------|
| Application   | Web/email request, DNS lookup, file upload              | Render HTML, show email, save file                    |
| Presentation  | Encrypt, compress, convert to ASCII/JSON               | Decrypt, decompress, translate to readable format     |
| Session       | Establish session, set dialog control                   | Accept session, maintain state, end session           |
| Transport     | Segment data, assign port, add seq#, TCP/UDP headers    | Reorder segments, remove duplicates, verify data      |
| Network       | Add source/destination IP, decide routing path          | Check IP, deliver to next hop or final destination    |
| Data Link     | Create frame, add MAC address, error detection          | Check MAC, validate frame, remove header              |
| Physical      | Convert bits to electrical/optical signal               | Convert signal back to bits, pass to Data Link layer  |

---

## Data Unit per OSI Layer

| OSI Layer    | Data Unit Name |
| ------------ | -------------- |
| Application  | Data           |
| Presentation | Data           |
| Session      | Data           |
| Transport    | Segment        |
| Network      | Packet         |
| Data Link    | Frame          |
| Physical     | Bits           |

---

## Devices Working at Each OSI Layer

| Layer        | Devices Involved            |
| ------------ | --------------------------- |
| Application  | Web browsers, email clients |
| Presentation | Encryption tools (SSL/TLS)  |
| Session      | Session managers, APIs      |
| Transport    | Firewalls, load balancers   |
| Network      | Routers                     |
| Data Link    | Switches, NICs              |
| Physical     | Cables, Hubs, Repeaters     |

---

## TCP/IP Model vs OSI Model

| OSI Model Layer         | TCP/IP Model Layer         |
|-------------------------|----------------------------|
| Application Layer       | Application Layer          |
| Presentation Layer      | Application Layer          |
| Session Layer           | Application Layer          |
| Transport Layer         | Transport Layer            |
| Network Layer           | Internet Layer             |
| Data Link Layer         | Network Access Layer       |
| Physical Layer          | Network Access Layer       |

---
## TCP/IP
![image](https://github.com/user-attachments/assets/0b7f3fc9-fa79-4ae6-a7dc-cb7bc9ff3447)

---
## TCP/IP Stack Details

| Layer                  | Protocols                              | Key Operations                                                                 | Input from Above Layer       | Processing Done                                | Output to Lower Layer             |
|------------------------|----------------------------------------|--------------------------------------------------------------------------------|-------------------------------|------------------------------------------------|------------------------------------|
| Link / Network Access  | Ethernet, Wi-Fi, PPP, ARP              | Frame creation, MAC addressing, media access, error detection, signaling       | IP Packet                    | Add MAC addresses, frame, error check, signal  | Bits on physical medium           |
| Internet               | IPv4, IPv6, ICMP, IGMP, ARP            | IP addressing, routing, fragmentation, error reporting                         | TCP/UDP Segment              | Add IP header, route, fragment                  | IP Packet to Link Layer           |
| Transport              | TCP, UDP, SCTP                         | Process delivery, reliability, sequencing, ports, error control                | Application data             | Add ports, TCP/UDP headers, manage ACKs         | Segment or Datagram to Internet   |
| Application            | HTTP, FTP, SMTP, DNS, SSH, DHCP       | User communication, data formatting, session handling                          | User/app-level data          | Format data, select protocol, create message    | Message/Data to Transport Layer   |

---

## HTTP Status Codes

| Code | Hex  | Category              | Mnemonic             | Description                                 |
|------|------|------------------------|----------------------|---------------------------------------------|
| 100  | 0x64 | 1xx - Informational    | Hold On              | Continue                                    |
| 101  | 0x65 |                        |                      | Switching Protocols                         |
| 102  | 0x66 |                        |                      | Processing (WebDAV)                         |
| 103  | 0x67 |                        |                      | Early Hints                                 |
| 200  | 0xC8 | 2xx - Success          | Here You Go          | OK                                          |
| 201  | 0xC9 |                        |                      | Created                                     |
| 202  | 0xCA |                        |                      | Accepted                                    |
| 203  | 0xCB |                        |                      | Non-Authoritative Information               |
| 204  | 0xCC |                        |                      | No Content                                  |
| 205  | 0xCD |                        |                      | Reset Content                               |
| 206  | 0xCE |                        |                      | Partial Content                             |
| 300  | 0x12C| 3xx - Redirection      | Go Elsewhere         | Multiple Choices                            |
| 301  | 0x12D|                        |                      | Moved Permanently                           |
| 302  | 0x12E|                        |                      | Found (Temporarily)                         |
| 303  | 0x12F|                        |                      | See Other                                   |
| 304  | 0x130|                        |                      | Not Modified                                |
| 305  | 0x131|                        |                      | Use Proxy                                   |
| 307  | 0x133|                        |                      | Temporary Redirect                          |
| 308  | 0x134|                        |                      | Permanent Redirect                          |
| 400  | 0x190| 4xx - Client Error     | You Messed Up        | Bad Request                                 |
| 401  | 0x191|                        |                      | Unauthorized                                |
| 402  | 0x192|                        |                      | Payment Required                            |
| 403  | 0x193|                        |                      | Forbidden                                   |
| 404  | 0x194|                        |                      | Not Found                                   |
| 405  | 0x195|                        |                      | Method Not Allowed                          |
| 406  | 0x196|                        |                      | Not Acceptable                              |
| 407  | 0x197|                        |                      | Proxy Authentication Required               |
| 408  | 0x198|                        |                      | Request Timeout                             |
| 409  | 0x199|                        |                      | Conflict                                    |
| 410  | 0x19A|                        |                      | Gone                                        |
| 418  | 0x1A2|                        | Joke Code            | I'm a teapot                                |
| 422  | 0x1A6|                        |                      | Unprocessable Entity                        |
| 429  | 0x1AD|                        |                      | Too Many Requests                           |
| 451  | 0x1C3|                        | Legal Block          | Unavailable For Legal Reasons               |
| 500  | 0x1F4| 5xx - Server Error     | I Messed Up          | Internal Server Error                       |
| 501  | 0x1F5|                        |                      | Not Implemented                             |
| 502  | 0x1F6|                        |                      | Bad Gateway                                 |
| 503  | 0x1F7|                        |                      | Service Unavailable                         |
| 504  | 0x1F8|                        |                      | Gateway Timeout                             |
| 507  | 0x1FB|                        |                      | Insufficient Storage                        |
| 511  | 0x1FF|                        |                      | Network Authentication Required             |

---

## HTTP Topics Overview

- **HTTP Methods**: GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD  
- **Status Codes**: 1xx - 5xx explained  
- **Headers**: Request/Response/Custom headers  
- **Cookies & Sessions**: Session tracking, user state  
- **HTTPS**: TLS handshake, certificates, encryption  
- **Caching**: Cache-Control, ETag, Expires  
- **Redirection**: 301 vs 302 handling  
- **Content Negotiation**: Accept, Content-Type, encoding  
- **Request Bodies**: Form data, JSON, URL params  
- **Performance**: HTTP keep-alive, HTTP/2, HTTP/3  
- **Testing Tools**: curl, Postman, Wireshark  
- **Proxy/Gateways**: Forward & Reverse proxies  
- **Connection Management**: TCP socket, timeouts  
- **Advanced Topics**: CORS, chunked encoding, preflight  
- **Real-World Usage**: REST APIs, GraphQL, webhooks  

---

## JWT Token Call Flow

```text
[Client] --> POST /login ------------------> [Server]
               credentials                  ↳ validate user
                                            ↳ generate JWT
                                            ↳ return JWT

[Client] <-- 200 OK + token ----------------

[Client] --> GET /dashboard ---------------> [Server]
               Authorization: Bearer token  ↳ Middleware
                                            ↳ verify token
                                            ↳ if valid: pass to handler
                                                      ↳ Controller returns data
                                            ↳ if invalid: return 401

[Client] <-- 200 OK or 401 Unauthorized <---
