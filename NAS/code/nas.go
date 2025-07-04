package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
	"net"
)

func main() {
	// Create output pcap file
	f, err := os.Create("dummy_nas_trace.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Initialize PCAP writer
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(65536, layers.LinkTypeEthernet)

	// Dummy NAS messages
	nasMessages := [][]byte{
		{0x7e, 0x00, 0x01, 0x01}, // Registration Request
		{0x7e, 0x00, 0x02, 0x01}, // Authentication Request
		{0x7e, 0x00, 0x02, 0x02}, // Authentication Response
		{0x7e, 0x00, 0x03, 0x01}, // Security Mode Command
		{0x7e, 0x00, 0x03, 0x02}, // Security Mode Complete
		{0x7e, 0x00, 0x04, 0x01}, // Registration Accept
		{0x7e, 0x00, 0x04, 0x02}, // Registration Complete
	}

	// Source and destination IPs
	srcIP := net.IP{192, 168, 1, 100}
	dstIP := net.IP{192, 168, 1, 200}

	for i, msg := range nasMessages {
		buf := gopacket.NewSerializeBuffer()
		opts := gopacket.SerializeOptions{FixLengths: true, ComputeChecksums: true}

		eth := layers.Ethernet{
			SrcMAC:       net.HardwareAddr{0x00, 0x0c, 0x29, 0x4f, 0x8e, 0x35},
			DstMAC:       net.HardwareAddr{0x00, 0x50, 0x56, 0xe6, 0x6b, 0x3d},
			EthernetType: layers.EthernetTypeIPv4,
		}

		ip := layers.IPv4{
			SrcIP:    srcIP,
			DstIP:    dstIP,
			Version:  4,
			TTL:      64,
			Protocol: layers.IPProtocolUDP,
		}

		udp := layers.UDP{
			SrcPort: 38412,
			DstPort: 38412,
		}
		udp.SetNetworkLayerForChecksum(&ip)

		payload := gopacket.Payload(msg)

		err := gopacket.SerializeLayers(buf, opts, &eth, &ip, &udp, payload)
		if err != nil {
			log.Fatal(err)
		}

		// Write packet with timestamp
		w.WritePacket(gopacket.CaptureInfo{
			Timestamp:     time.Now().Add(time.Duration(i) * time.Second),
			CaptureLength: len(buf.Bytes()),
			Length:        len(buf.Bytes()),
		}, buf.Bytes())
	}

	fmt.Println("✅ NAS .pcap file created: dummy_nas_trace.pcap")
}
