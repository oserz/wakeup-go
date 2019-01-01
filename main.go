package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
)

type MACAddress [6]byte

type MagicPacket struct {
	header  [6]byte
	payload [16]MACAddress
}

var (
	delims = ":-"
	reMAC  = regexp.MustCompile(`^([0-9a-fA-F]{2}[` + delims + `]){5}([0-9a-fA-F]{2})$`)

	macAddr        string
	bcastInterface string
)

func init() {
	flag.StringVar(&macAddr, "mac", "", "wake up by MAC address, a valid MAC is like: 01-02-03-04-05-06, 01:02:03:04:05:06")
	flag.StringVar(&bcastInterface, "interface", "", "outbound interface to broadcast using")
	flag.Usage = usage
}

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func NewMagicBuff(mac string) (*MagicPacket, error) {
	var packet MagicPacket
	var macAddr MACAddress

	if !reMAC.MatchString(mac) {
		return nil, fmt.Errorf("%s is not a IEEE 802 MAC-48 address", mac)
	}

	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		return nil, err
	}

	for idx := range macAddr {
		macAddr[idx] = hwAddr[idx]
	}

	for idx := range packet.header {
		packet.header[idx] = 0xFF
	}

	for idx := range packet.payload {
		packet.payload[idx] = macAddr
	}

	return &packet, nil
}

func main() {
	var localaddr *net.UDPAddr
	flag.Parse()

	if macAddr == "" {
		usage()
	}

	fmt.Printf("Mac is %s, interface is %s \n", macAddr, bcastInterface)
	ief, _ := net.InterfaceByName(bcastInterface)
	addrs, _ := ief.Addrs()
	for _, addr := range addrs {
		switch ip := addr.(type) {
		case *net.IPNet:
			if ip.IP.DefaultMask() != nil {
				fmt.Println(ip.IP)
				localaddr = &net.UDPAddr{
					IP: ip.IP,
				}
			}
		}
	}

	//	ipwake := "192.168.2.15"
	//	addrip := net.ParseIP(ipwake)

	remoteAddr, _ := net.ResolveUDPAddr("udp", "255.255.255.255:9")

	conn, _ := net.DialUDP("udp", localaddr, remoteAddr)

	defer conn.Close()
	var mp *MagicPacket

	mp, err := NewMagicBuff(macAddr)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	if err = binary.Write(&buf, binary.BigEndian, mp); err != nil {
		return
	}

	fmt.Printf("Attempting to send a magic packet to MAC %s, local IP: %s \n", macAddr, localaddr.IP)
	conn.Write(buf.Bytes())

	fmt.Println("Done!")
}
