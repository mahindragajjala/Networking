package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := "hello server"
	conn.Write([]byte(message))
	fmt.Println("Message sent.")
}
                        You are not writing the TCP/IP protocol manually, but you’re using Go’s net package, which:
                        
                         Your Code Calls       Underlying Action                                             
                         `net.Dial()`          Uses OS system call → creates TCP socket (`socket()`)         
                         `conn.Write([]byte)`  Sends user data to TCP buffer → OS adds TCP/IP headers        
                         OS Kernel             Constructs: Ethernet + IP + TCP headers + Data                
                         NIC / Loopback        Sends complete frame over network (or loopback for 127.0.0.1) 

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
	// Open device
	handle, err := pcap.OpenLive("lo", 65535, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Construct Ethernet (not used for loopback), IP, TCP layers
	ip := &layers.IPv4{
		SrcIP:    net.ParseIP("127.0.0.1"),
		DstIP:    net.ParseIP("127.0.0.1"),
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}

	tcp := &layers.TCP{
		SrcPort: layers.TCPPort(42498),
		DstPort: layers.TCPPort(8080),
		SYN:     true,
		Seq:     110,
	}
	tcp.SetNetworkLayerForChecksum(ip)

	// Add payload
	payload := []byte("hello server")

	// Serialize
	buffer := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	err = gopacket.SerializeLayers(buffer, opts,
		ip,
		tcp,
		gopacket.Payload(payload),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Send packet
	err = handle.WritePacketData(buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Raw TCP/IP packet sent.")
	time.Sleep(time.Second)
}
