package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"unicode"

	"github.com/thecsw/darkness/emilia"
)

// clean cleans up the directory
func clean() {
	cleanCmd := flag.NewFlagSet("clean", flag.ExitOnError)
	cleanCmd.StringVar(&workDir, "dir", ".", "where do I look for files")
	cleanCmd.StringVar(&darknessToml, "conf", "darkness.toml", "location of darkness.toml")
	cleanCmd.Parse(os.Args[2:])
	emilia.InitDarkness(darknessToml)
	orgfiles, err := findFilesByExt(workDir, emilia.Config.Project.Input)
	if err != nil {
		fmt.Printf("failed to find files by extension %s: %s",
			emilia.Config.Project.Input, err.Error())
		os.Exit(1)
	}
	for _, orgfile := range orgfiles {
		toRemove := getTarget(orgfile)
		if err := os.Remove(toRemove); err != nil && !os.IsNotExist(err) {
			fmt.Println(toRemove, "failed to delete: "+err.Error())
		}
	}
}

// megumin blows up the directory
func megumin() {
	explosionCmd := flag.NewFlagSet("megumin", flag.ExitOnError)
	explosionCmd.StringVar(&workDir, "dir", ".", "where do I look for files")
	explosionCmd.StringVar(&darknessToml, "conf", "darkness.toml", "location of darkness.toml")
	explosionCmd.Parse(os.Args[2:])
	emilia.InitDarkness(darknessToml)

	orgfiles, err := findFilesByExt(workDir, emilia.Config.Project.Input)
	if err != nil {
		fmt.Printf("failed to find files by extension %s: %s",
			emilia.Config.Project.Input, err.Error())
		os.Exit(1)
	}
	delayedLinesPrint([]string{
		"Darker than black, darker than darkness, combine with my intense crimson.",
		"Time to wake up, descend to these borders and appear as an intangible distortion.",
		"Dance, dance, dance!",
		"May a destructive force flood my torrent of power, a destructive force like no other!",
		"Send all creation to its source!",
		"Come out of your abyss!",
		"Humanity knows no other more powerful offensive technique!",
		"It is the ultimate magical attack!",
		"Explosion!",
	})
	for _, orgfile := range orgfiles {
		toRemove := getTarget(orgfile)
		if err := os.Remove(toRemove); err != nil && !os.IsNotExist(err) {
			fmt.Println(toRemove, "failed to blow up!!")
		}
		fmt.Println(toRemove, "went boom!")
		time.Sleep(50 * time.Millisecond)
	}
	delayedLinesPrint([]string{
		"Wahahahahaha!",
		"My name is Megumin, the number one mage of Axel!",
		"Come, you shall all become my experience points today!",
	})
}

// delayedLinesPrint prints lines with a delay
func delayedLinesPrint(lines []string) {
	for _, line := range lines {
		time.Sleep(200 * time.Millisecond)
		delayedSentencePrint(line)
		time.Sleep(900 * time.Millisecond)
		fmt.Printf("\n")
	}
}

// delayedSentencePrint prints a sentence with a delay
func delayedSentencePrint(line string) {
	for _, c := range line {
		fmt.Printf("%c", c)
		time.Sleep(60 * time.Millisecond)
		if unicode.IsPunct(c) {
			time.Sleep(400 * time.Millisecond)
		}
	}
}