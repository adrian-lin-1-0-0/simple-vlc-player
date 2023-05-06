package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/adrian-lin-1-0-0/simple-vlc-player/vlc"
)

func clearScreen() {
	cmd := exec.Command("clear") //Linux
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showHelp() {
	fmt.Println(`
   [↑] : to increase volume
   [↓] : to decrease volume
   [→] : to forward 10 seconds
   [←] : to backward 10 seconds
   [q] : to exit

  `)
}

func showAppInfo(player *vlc.Player) {
	clearScreen()
	showHelp()
	showVolumeBar(player)
	fmt.Print(" ")
	showMediaTime(player)

}

var volumeBars = []string{" ", "▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "█", "█"}

func buildVolumeBar(volume int) string {
	return volumeBars[volume/10]
}

func showVolumeBar(player *vlc.Player) {
	fmt.Print(buildVolumeBar(player.GetVolume()))
}

func buildProgressBar(current, total int) string {
	const width = 40
	completed := int(float64(current) / float64(total) * width)
	bar := strings.Repeat("█", completed) + strings.Repeat("░", width-completed)
	return fmt.Sprintf("%s %02d:%02d %02d:%02d\n", bar, current/60, current%60, total/60, total%60)
}

func showMediaTime(player *vlc.Player) {
	fmt.Println(buildProgressBar(player.GetMediaTime()/1000, player.GetMediaLength()/1000))
}
