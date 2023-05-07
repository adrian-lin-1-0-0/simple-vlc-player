package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/adrian-lin-1-0-0/simple-vlc-player/utils"
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

	showHelp := true

	go func() {
		for {
			time.Sleep(1 * time.Second)
			utils.ShowAppInfo(player, &showHelp)
		}
	}()

	for {
		utils.ShowAppInfo(player, &showHelp)
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyArrowUp {
			utils.VolumeUp10(player)
		}

		if key == keyboard.KeyArrowDown {
			utils.VolumeDown10(player)
		}

		if key == keyboard.KeyArrowRight {
			utils.Forward10Second(player)
		}

		if key == keyboard.KeyArrowLeft {
			utils.Backward10Second(player)
		}

		if char == 'h' {
			showHelp = !showHelp
		}

		if char == 'q' {
			break
		}
	}

	os.Exit(0)
}
