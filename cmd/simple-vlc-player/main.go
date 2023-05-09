package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/adrian-lin-1-0-0/simple-vlc-player/utils"
	"github.com/eiannone/keyboard"
)

func main() {
	fileName := flag.String("f", "", "File name")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Please specify a file name using the -f option")
		return
	}

	mediaPlayer := utils.NewMediaPlayer(fileName)
	defer mediaPlayer.Release()

	err := mediaPlayer.Play()
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
			utils.ShowAppInfo(mediaPlayer, &showHelp)
		}
	}()

	for {
		utils.ShowAppInfo(mediaPlayer, &showHelp)
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			mediaPlayer.VolumeUp10()
		case keyboard.KeyArrowDown:
			mediaPlayer.VolumeDown10()
		case keyboard.KeyArrowRight:
			mediaPlayer.Forward10Second()
		case keyboard.KeyArrowLeft:
			mediaPlayer.Backward10Second()
		case 0:
			switch char {
			case 'h':
				showHelp = !showHelp
			case 'q':
				return
			}
		}
	}
}
