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

## TCP/IP
![image](https://github.com/user-attachments/assets/0b7f3fc9-fa79-4ae6-a7dc-cb7bc9ff3447)

| 🧱 Layer (Bottom → Top)      | 📜 Protocols (Examples)                         | ⚙️ Key Operations in Layer                                                                                             | 🔽 Input from Above Layer               | 🔄 Processing Done in the Layer                                                                                                        | 🔼 Output to Lower Layer                    |
| ---------------------------- | ----------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | --------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| **1. Link / Network Access** | Ethernet (IEEE 802.3), Wi-Fi (802.11), PPP, ARP | - Frame creation<br>- MAC addressing<br>- Media access (CSMA/CD)<br>- Error detection (CRC)<br>- Physical signaling    | IP Packet                               | ➤ Add MAC addresses<br>➤ Wrap in Ethernet frame<br>➤ Perform error check (CRC)<br>➤ Transmit as bits on physical medium                | Bits sent over wire or wireless media       |
| **2. Internet**              | IPv4, IPv6, ICMP, IGMP, ARP                     | - IP addressing<br>- Routing<br>- Packet fragmentation/reassembly<br>- Error reporting (ICMP)                          | TCP/UDP Segment                         | ➤ Add IP header (source/destination IP)<br>➤ Check TTL<br>➤ Route via routers<br>➤ Fragment if needed                                  | IP Packet to Link Layer                     |
| **3. Transport**             | TCP, UDP, SCTP                                  | - Process-to-process delivery<br>- Reliability (TCP)<br>- Flow control<br>- Error detection<br>- Sequencing<br>- Ports | Application data (e.g., HTTP request)   | ➤ Add source/destination ports<br>➤ Add TCP/UDP header<br>➤ Manage ACKs, retransmissions (TCP)<br>➤ Optional flow & congestion control | Segment (TCP) or Datagram (UDP) to Internet |
| **4. Application**           | HTTP/HTTPS, FTP, SMTP, DNS, SSH, Telnet, DHCP   | - User-level communication<br>- Protocol execution<br>- Session handling<br>- Data formatting<br>- Service advertising | User input or application-level message | ➤ Format/encode data (HTML, JSON, etc.)<br>➤ Choose protocol (e.g., HTTP)<br>➤ Create message for delivery                             | Message/Data to Transport Layer             |

## HTTP - Hypertext Transfer Protocol

# 📦 HTTP Status Code

| Code | Hex  | Category              | Mnemonic             | Description                                 |
|------|------|------------------------|----------------------|---------------------------------------------|
| 100  | 0x64 | 1xx - Informational    | 🕐 Hold On           | Continue                                    |
| 101  | 0x65 |                        |                      | Switching Protocols                         |
| 102  | 0x66 |                        |                      | Processing (WebDAV)                         |
| 103  | 0x67 |                        |                      | Early Hints                                 |
| 200  | 0xC8 | 2xx - Success          | ✅ Here You Go       | OK                                          |
| 201  | 0xC9 |                        |                      | Created                                     |
| 202  | 0xCA |                        |                      | Accepted                                    |
| 203  | 0xCB |                        |                      | Non-Authoritative Information               |
| 204  | 0xCC |                        |                      | No Content                                  |
| 205  | 0xCD |                        |                      | Reset Content                               |
| 206  | 0xCE |                        |                      | Partial Content                             |
| 300  | 0x12C| 3xx - Redirection      | 🔀 Go Elsewhere      | Multiple Choices                            |
| 301  | 0x12D|                        |                      | Moved Permanently                           |
| 302  | 0x12E|                        |                      | Found (Temporarily)                         |
| 303  | 0x12F|                        |                      | See Other                                   |
| 304  | 0x130|                        |                      | Not Modified                                |
| 305  | 0x131|                        |                      | Use Proxy                                   |
| 307  | 0x133|                        |                      | Temporary Redirect                          |
| 308  | 0x134|                        |                      | Permanent Redirect                          |
| 400  | 0x190| 4xx - Client Error     | ❌ You Messed Up     | Bad Request                                 |
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
| 418  | 0x1A2|                        | ☕ Joke Code          | I'm a teapot                                |
| 422  | 0x1A6|                        |                      | Unprocessable Entity                        |
| 429  | 0x1AD|                        |                      | Too Many Requests                           |
| 451  | 0x1C3|                        | 📛 Legal Block       | Unavailable For Legal Reasons               |
| 500  | 0x1F4| 5xx - Server Error     | 💥 I Messed Up       | Internal Server Error                       |
| 501  | 0x1F5|                        |                      | Not Implemented                             |
| 502  | 0x1F6|                        |                      | Bad Gateway                                 |
| 503  | 0x1F7|                        |                      | Service Unavailable                         |
| 504  | 0x1F8|                        |                      | Gateway Timeout                             |
| 507  | 0x1FB|                        |                      | Insufficient Storage                        |
| 511  | 0x1FF|                        |                      | Network Authentication Required             |

## HTTP - TOPICS
 📑 HTTP Methods           GET, POST, PUT, DELETE, PATCH, OPTIONS,                              HEAD               
 ⚙️ Status Codes           1xx to 5xx – what they mean                                
 📡 Headers                Request headers, response headers, custom headers          
 🍪 Cookies & Sessions     How cookies work, session tracking, authentication         
 🔒 HTTPS                  TLS handshake, certificate, encryption                     
 🔄 Caching                Cache-Control, ETag, Expires                               
 ↪️ Redirection            301 vs 302, how browsers follow redirects                  
 📶 Content Negotiation    Accept, Accept-Encoding, Content-Type                      
 📤 Request Body/Query     Form data, JSON payloads, URL parameters                   
 📈 Performance            Keep-alive, pipelining, HTTP/2/3 improvements              
 🧪 Testing HTTP           Using `curl`, `Postman`, `Wireshark` to inspect traffic    
 🔀 Proxy & Gateway        Reverse proxies (nginx), forward proxies                   
 🚪 Connection Management  TCP socket handling, timeout, keep-alive                   
 🧵 Advanced Topics        Chunked transfer encoding, CORS, preflight requests        
 📊 Real-World Usage       REST APIs, GraphQL over HTTP, webhooks  


## JWT TOKEN 
                     [Client] --> POST /login ------------------> [Server]
                                       credentials               ↳ validate user
                                                                 ↳ generate JWT
                                                                 ↳ return JWT
                     
                     [Client] <-- 200 OK + token ---------------
                     
                     [Client] --> GET /dashboard ---------------> [Server]
                                    Authorization: Bearer token ↳ Middleware
                                                                 ↳ verify token
                                                                 ↳ if valid: pass to handler
                                                                        ↳ Controller returns data
                                                                 ↳ if invalid: return 401
                     
                     [Client] <-- 200 OK or 401 Unauthorized <---
