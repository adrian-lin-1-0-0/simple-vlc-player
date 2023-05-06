package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/adrian-lin-1-0-0/simple-vlc-player/vlc"
	"github.com/eiannone/keyboard"
)

func main() {
	fileName := flag.String("f", "", "File name")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Please specify a file name using the -f option")
		return
	}

	vlcInstance, err := vlc.Init("--no-video", "--quiet")
	if err != nil {
		panic(err)
	}
	defer vlcInstance.Release()

	media, err := vlcInstance.NewMedia(*fileName)
	if err != nil {
		panic(err)
	}
	defer media.Release()

	player, err := vlcInstance.NewPlayer()
	if err != nil {
		panic(err)
	}
	defer player.Release()

	err = player.SetMedia(media)
	if err != nil {
		panic(err)
	}

	err = player.Play()
	if err != nil {
		panic(err)
	}

	if err = keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	//wait for vlc to play the media and ready to prepare buffer
	time.Sleep(1 * time.Second)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			showAppInfo(player)
		}
	}()

	for {
		showAppInfo(player)
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyArrowUp {
			volumeUp10(player)
		}

		if key == keyboard.KeyArrowDown {
			volumeDown10(player)
		}

		if key == keyboard.KeyArrowRight {
			forward10Second(player)
		}

		if key == keyboard.KeyArrowLeft {
			backward10Second(player)
		}

		if char == 'q' {
			break
		}
	}

	os.Exit(0)
}
