package main

import (
	"fmt"
	"playground/20072018/CLI_emoji_Chooser/emoji"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"
)

func main() {

	// Will need expanding upon above the beginning
	faces := emoji.MakeFaceGroup()
	people := emoji.MakePeopleGroup()
	el := emoji.EmojiList{[]emoji.Group{faces, people}}

	selectGroup := promptui.Select{
		Label: "Group (Faces are default)",
		Items: el.GetGroupNames(),
	}

	groupNo, group, err := selectGroup.Run()

	if err != nil {
		fmt.Println(err)
	}

	if group == "" {
		group = "face"
	}

	var items []string

	for sg := range el.Groups[groupNo].SubGroups {
		if el.Groups[groupNo].SubGroups[sg].Group == group {
			items = append(items, el.Groups[groupNo].SubGroups[sg].SubGroupNames[0])
		}
	}

	selectSub := promptui.Select{
		Label: "Select Emoji catagory",
		Items: items,
	}

	subGNo, _, err := selectSub.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var emojis []string
	for e := range el.Groups[groupNo].SubGroups[subGNo].Emojis {
		em := el.Groups[groupNo].SubGroups[subGNo].Emojis[e]
		emojis = append(emojis, " => "+em.Symbol+" : \""+strings.Join(em.Names, "\"")+"\"")
	}

	selectEmoji := promptui.Select{
		Label: " Select \"" + strings.Join(el.Groups[groupNo].SubGroups[subGNo].SubGroupNames, "\" , \"") + "\" Emoji",
		Items: emojis,
	}

	emojiNo, _, err := selectEmoji.Run()

	clipboard.WriteAll(el.Groups[groupNo].SubGroups[subGNo].Emojis[emojiNo].Symbol)

	fmt.Println(groupNo)
	fmt.Println(subGNo)
	fmt.Println(emojiNo)
	fmt.Printf("You choose %q\n", el.Groups[groupNo].SubGroups[subGNo].Emojis[emojiNo].Symbol+" - Now copied to clipboard")
}
