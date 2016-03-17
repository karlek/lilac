package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Narsil/alsa-go"
	"github.com/eaburns/flac"
)

func main() {
	if err := play(); err != nil {
		log.Fatal(err)
	}
}

func play() (err error) {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalln("lilac *.flac")
	}
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		return err
	}
	dec, err := flac.NewDecoder(f)
	if err != nil {
		return err
	}
	h := alsa.New()
	defer h.Close()
	err = h.Open("default", alsa.StreamTypePlayback, alsa.ModeBlock)
	if err != nil {
		fmt.Printf("Open failed. %s", err)
	}

	h.SampleFormat = alsa.SampleFormatS16LE
	h.SampleRate = dec.SampleRate
	h.Channels = dec.NChannels
	err = h.ApplyHwParams()
	if err != nil {
		fmt.Printf("SetHwParams failed. %s", err)
	}

	for {
		buf, err := dec.Next()
		if err != nil {
			return err
		}
		err = binary.Write(h, binary.BigEndian, buf)
		if err != nil {
			return err
		}
	}
	if err != nil {
		fmt.Printf("Write failed %s", err)
	}
	return nil
}
