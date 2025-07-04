![image](https://github.com/user-attachments/assets/527600cc-a3e1-4a54-8701-a48e6477926e)

![image](https://github.com/user-attachments/assets/b4b35449-e507-4971-b192-de430f310c98)


![image](https://github.com/user-attachments/assets/f578f6a6-49f2-451a-a5b2-d4cfc0697a15)



![image](https://github.com/user-attachments/assets/b5aa8551-8612-4e48-b84e-1efcd9750f49)


Protocol Layers(Top to Bottom):
| Layer                 | Function                          | Protocols           |
| --------------------- | --------------------------------- | ------------------- |
| **Application Layer** | Network services and applications | NAS (for CP)        |
| **Transport Layer**   | Reliable transmission             | SCTP (CP), UDP (UP) |
| **Network Layer**     | IP routing                        | IP                  |
| **Data Link Layer**   | Framing and MAC addressing        | PDCP, RLC, MAC      |
| **Physical Layer**    | Bit transmission over air         | PHY                 |


Control Plane Protocol Stack (Signaling Path)
NAS ↔ NGAP ↔ SCTP ↔ IP ↔ PDCP ↔ RLC ↔ MAC ↔ PHY

| Layer    | Component                            | Description                                                                    |
| -------- | ------------------------------------ | ------------------------------------------------------------------------------ |
| **NAS**  | Non-Access Stratum                   | Between UE and 5GC; used for authentication, registration, session management. |
| **NGAP** | Next Generation Application Protocol | Between gNB and AMF; handles UE context, mobility, signaling.                  |
| **SCTP** | Stream Control Transmission Protocol | Transport of NGAP; reliable transport protocol.                                |
| **IP**   | Internet Protocol                    | Routing packets.                                                               |
| **PDCP** | Packet Data Convergence Protocol     | Header compression, ciphering, integrity protection.                           |
| **RLC**  | Radio Link Control                   | Segmentation, reassembly, retransmission.                                      |
| **MAC**  | Medium Access Control                | Scheduling, multiplexing of logical channels.                                  |
| **PHY**  | Physical Layer                       | Modulation, coding, transmission over radio.                                   |


User Plane Protocol Stack (Data Transfer Path)
GTP-U ↔ UDP ↔ IP ↔ PDCP ↔ RLC ↔ MAC ↔ PHY

| Layer     | Component                            | Description                                |
| --------- | ------------------------------------ | ------------------------------------------ |
| **GTP-U** | GPRS Tunneling Protocol - User Plane | Tunnels user data from gNB to UPF in core. |
| **UDP**   | User Datagram Protocol               | Lightweight transport; no reliability.     |
| **IP**    | Internet Protocol                    | Used for routing.                          |
| **PDCP**  | Same as above.                       |                                            |
| **RLC**   | Same as above.                       |                                            |
| **MAC**   | Same as above.                       |                                            |
| **PHY**   | Same as above.                       |                                            |

Core Network Protocols
| Core Function | Protocol                                             | Purpose                      |
| ------------- | ---------------------------------------------------- | ---------------------------- |
| AMF ↔ UE      | NAS                                                  | Registration, authentication |
| AMF ↔ gNB     | NGAP                                                 | Mobility, context mgmt       |
| SMF ↔ UPF     | PFCP                                                 | Session mgmt, QoS            |
| UPF ↔ gNB     | GTP-U                                                | User data forwarding         |
| UDM, AUSF     | HTTP/2 (REST API over Service-Based Interface - SBI) | Core internal communication  |

