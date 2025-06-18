package main

import (
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Open loopback interface for sending
	handle, err := pcap.OpenLive("lo", 65535, false, pcap.BlockForever)
	if err != nil {
		log.Fatal("❌ Error opening interface:", err)
	}
	defer handle.Close()

	srcIP := net.ParseIP("127.0.0.1")
	dstIP := net.ParseIP("127.0.0.1")

	ip := &layers.IPv4{
		Version:  4,
		IHL:      5,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
		SrcIP:    srcIP,
		DstIP:    dstIP,
	}

	tcp := &layers.TCP{
		SrcPort: 42498,
		DstPort: 8080,
		Seq:     100,
		ACK:     true,
		PSH:     true,
	}
	tcp.SetNetworkLayerForChecksum(ip)

	payload := []byte("hello server")

	// Serialize layers
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	/* err = gopacket.SerializeLayers(buf, opts,
		ip,
		tcp,
		gopacket.Payload(payload),

	) */

	err = gopacket.SerializeLayers(buf, opts,
		ip,                        // Internet Layer
		tcp,                       // Transport Layer
		gopacket.Payload(payload), // Application Layer
	)

	if err != nil {
		log.Fatal("❌ Serialization error:", err)
	}

	err = handle.WritePacketData(buf.Bytes())
	if err != nil {
		log.Fatal("❌ Packet sending error:", err)
	}

	log.Println("✅ Raw TCP packet with payload 'hello server' sent.")
	time.Sleep(time.Second)
}

/*
[Application Layer]     → "hello server"
        ↓
[Transport Layer]       → TCP Header (SrcPort, DstPort, Seq, PSH, ACK)
        ↓
[Internet Layer]        → IPv4 Header (SrcIP, DstIP, TTL, Protocol)
        ↓
[Link Layer]            → Not explicitly added in Go (PCAP handles it)

*/
