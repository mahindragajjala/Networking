# 5G UE Registration Call Flow (ASCII)

+--------+                +--------+                 +--------+
|   UE   |                |  RAN   |                 |  AMF   |
+--------+                +--------+                 +--------+
     |                         |                          |
     | 1. Power ON             |                          |
     |------------------------>|                          |
     |                         |                          |
     | 2. RRC Connection Req   |                          |
     |------------------------>|                          |
     |                         |                          |
     | 3. RRC Setup            |                          |
     |<------------------------|                          |
     |                         |                          |
     | 4. RRC Setup Complete + NAS Registration Request   |
     |------------------------>|                          |
     |                         |                          |
     |                         | 5. NGAP Initial UE Msg   |
     |                         |    + NAS Registration Req|
     |                         |------------------------->|
     |                         |                          |
     |                         |     6. NAS Auth Request  |
     |                         |<-------------------------|
     |                         |                          |
     | 7. RRC DL Info Transfer + NAS Auth Request         |
     |<------------------------|                          |
     |                         |                          |
     | 8. NAS Auth Response    |                          |
     |------------------------>|                          |
     |                         |                          |
     |                         | 9. NGAP Uplink NAS Msg   |
     |                         |    + NAS Auth Response   |
     |                         |------------------------->|
     |                         |                          |
     |                         |     10. NAS Sec Mode Cmd |
     |                         |<-------------------------|
     |                         |                          |
     | 11. RRC DL Info Transfer + NAS Sec Mode Command    |
     |<------------------------|                          |
     |                         |                          |
     | 12. NAS Sec Mode Complete                          |
     |------------------------>|                          |
     |                         |                          |
     |                         | 13. NGAP Uplink NAS Msg  |
     |                         |     + Sec Mode Complete  |
     |                         |------------------------->|
     |                         |                          |
     |                         |     14. NAS Reg Accept   |
     |                         |<-------------------------|
     |                         |                          |
     | 15. RRC DL Info Transfer + NAS Registration Accept |
     |<------------------------|                          |
     |                         |                          |
     | 16. NAS Registration Complete                     |
     |------------------------>|                          |
     |                         |                          |
     |                         | 17. NGAP Uplink NAS Msg  |
     |                         |     + Reg Complete       |
     |                         |------------------------->|
     |                         |                          |
     |                         | 18. NGAP Initial Ctx Setup Req (AMF → RAN)  
     |                         |------------------------->|
     |                         |                          |
     | 19. RRC Reconfig        |                          |
     |<------------------------|                          |
     |                         |                          |
     | 20. RRC Reconfig Complete                          |
     |------------------------>|                          |
     |                         |                          |
     |                         | 21. Initial Ctx Setup Rsp|
     |                         |------------------------->|
     |                         |                          |
     |<====== UE Registered, PDU Session Setup Can Follow ======>|


# 📡 NAS Protocol During Initial Registration (UE ⇄ RAN ⇄ AMF)

This document explains how the **NAS (Non-Access Stratum) protocol** functions during the **initial registration procedure** in a 5G network between the **UE (User Equipment)** and the **AMF (Access and Mobility Management Function)**, with the **RAN (gNB)** acting as a transparent forwarder.

---

## 🔄 Step-by-Step Flow: Initial NAS Registration

### 1. UE → RAN (gNB): RRC + NAS

- UE sends:
  - `RRC Connection Request` to establish a connection with gNB.
  - After RRC connection is accepted, it sends `RRC Connection Setup Complete`, **carrying a NAS: Registration Request**.
- ✅ **NAS messages are encapsulated inside RRC messages.**

---

### 2. RAN (gNB) → AMF: NGAP + NAS

- gNB extracts the NAS Registration Request.
- It wraps the NAS message inside an `NGAP Initial UE Message`.
- Forwards it to AMF via the **N2 interface (SCTP protocol)**.
- 🧠 gNB does **not interpret NAS** — it just **relays** NAS transparently.

---

### 3. AMF ↔ UE: NAS Signaling Exchanges (via gNB)

After AMF receives the registration request, it continues the NAS signaling procedure:

- `NAS Authentication Request / Response`
- `Security Mode Command / Complete`
- `Registration Accept / Complete`

These NAS messages are:
- Generated by AMF
- Sent to UE via gNB
- **Encapsulated inside NGAP + RRC messages**

---

## 🧭 Protocol Stack Overview (UE ⇄ AMF via gNB)

| Layer         | UE                       | gNB                        | AMF                      |
|---------------|---------------------------|----------------------------|--------------------------|
| Application   | NAS                       | (Transparent to NAS)       | NAS                      |
| Signaling     | RRC                       | NGAP                       | NGAP                     |
| Transport     | PDCP/RLC/MAC/PHY (5G NR) | SCTP over IP               | SCTP over IP             |

