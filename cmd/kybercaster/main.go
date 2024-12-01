package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/storskegg/poc-kybercaster/udp"
)

const appName = "kybercaster"

var cmdRoot = cobra.Command{
	Use:     appName,
	Short:   "listens and multicasts messages with very strong peer encryption",
	Version: "v0.0.1",
	Run:     execRoot,
}

func init() {
	cmdRoot.Flags().IPVar(&flagIPv4, "ip", net.ParseIP(udp.DefaultMulticastIPv4), "IPv4")
	cmdRoot.Flags().StringVar(&flagPort, "port", udp.DefaultMulticastPort, "port")
	cmdRoot.Flags().StringVar(&flagFrom, "from", "foo", "a unique string identifier")
	cmdRoot.Flags().DurationVar(&flagCadenceGeneral, "cadence", 5*time.Second, "broadcast cadence for general data (format: 1h2m3s)")
	cmdRoot.Flags().DurationVar(&flagCadencePubkey, "cadence-pubkey", 15*time.Second, "broadcast cadence for public keys")

	if err := cmdRoot.MarkFlagRequired("from"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	flagIPv4           net.IP
	flagPort           string
	flagFrom           string
	flagCadenceGeneral time.Duration
	flagCadencePubkey  time.Duration
)

var localAddr string

func execRoot(cmd *cobra.Command, args []string) {
	if flagIPv4.To4() == nil {
		cmd.Usage()
		return
	}

	addr := flagIPv4.String() + ":" + flagPort

	chanDone := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	udp.Listen(addr, handler)

	conn, err := udp.NewCaster(addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr = conn.LocalAddr().String()
	fmt.Println("Local Address:", localAddr)

	go func() {
		ticker := time.NewTicker(flagCadenceGeneral)
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				conn.Write([]byte("hello from " + flagFrom))
			}
		}
	}()

	<-chanDone
}

func handler(src *net.UDPAddr, n int, b []byte) {
	action := "RECV"
	if src.String() == localAddr {
		action = "SELF"
	}
	fmt.Printf("%s: %d bytes read from %s\n", action, n, src.String())
	fmt.Printf("---- '%s' at %s\n", string(b[:n]), time.Now().Format(time.RFC3339))
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("\nErr: recovered panic")
			fmt.Println(err)
		}
	}()

	if err := cmdRoot.Execute(); err != nil {
		panic(err)
	}
}
