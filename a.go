package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/Narsil/alsa-go"
	"github.com/mewkiz/flac"
)

func main() {
	if err := play(); err != nil {
		log.Fatal(err)
	}
}

func play() (err error) {
	s, err := flac.Open("/home/_/go/src/github.com/karlek/lolfi/a.flac")
	if err != nil {
		return err
	}
	defer s.Close()
	h := alsa.New()
	defer h.Close()
	err = h.Open("teewav", alsa.StreamTypePlayback, alsa.ModeBlock)
	if err != nil {
		fmt.Printf("Open failed. %s", err)
	}

	h.SampleFormat = alsa.SampleFormatU8
	h.SampleRate = int(s.Info.SampleRate)
	h.Channels = int(s.Info.NChannels)
	err = h.ApplyHwParams()
	if err != nil {
		fmt.Printf("SetHwParams failed. %s", err)
	}
	// buf := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for {
		f, err := s.ParseNext()
		if err != nil {
			return err
		}
		frames := f.Subframes
		for _, frame := range frames {
			binary.Write(h, binary.BigEndian, frame.Samples)
		}
	}
	// buf = []byte(frames[0].Samples)

	// n, err := h.Write(buf)
	if err != nil {
		fmt.Printf("Write failed %s", err)
	}
	// if n != len(buf) {
	// 	fmt.Printf("Did not write all data (Wrote %d, expected %d)", n, len(buf))
	// }
	return nil
}
