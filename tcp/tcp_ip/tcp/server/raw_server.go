package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Open the loopback interface for capturing
	handle, err := pcap.OpenLive("lo", 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("❌ Error opening interface:", err)
	}
	defer handle.Close()

	fmt.Println("🟢 Server is listening for TCP packets on port 8080...")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		tcpLayer := packet.Layer(layers.LayerTypeTCP)

		if ipLayer != nil && tcpLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			tcp, _ := tcpLayer.(*layers.TCP)

			// Only accept packets destined to port 8080 with payload
			if tcp.DstPort == 8080 && len(tcp.Payload) > 0 {
				fmt.Println("📦 TCP Packet Received:")
				fmt.Printf("➡️  From %s:%d\n", ip.SrcIP, tcp.SrcPort)
				fmt.Printf("⬅️  To   %s:%d\n", ip.DstIP, tcp.DstPort)
				fmt.Printf("📝 Payload: %s\n", string(tcp.Payload))
				break
			}
		}
	}
}
