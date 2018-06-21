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
	sports := emoji.MakeSportGroup()
	family := emoji.MakeFamilyGroup()
	el := emoji.EmojiList{[]emoji.Group{faces, people, sports, family}}

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
	emojiChoice := el.Groups[groupNo].SubGroups[subGNo].Emojis

	searcher := func(input string, index int) bool {
		emoji := emojiChoice[index]
		name := strings.Replace(strings.ToLower(emoji.Names[0]), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   " => {{ .Names[0] | cyan }} ({{ .Symbol | red }})",
		Inactive: "  {{ .Names[0] | cyan }} ({{ .Symbol | red }})",
		Selected: "  => {{ .Names[0] | red | cyan }}",
		Details: `
--------- Choice ----------
{{ "Name:" | faint }}	{{ .Names[0] }}
{{ "Symbol:" | faint }}	{{ .Symbol }}`,
	}

	selectEmoji := promptui.Select{
		Label:     " Select \"" + strings.Join(el.Groups[groupNo].SubGroups[subGNo].SubGroupNames, "\" , \"") + "\" Emoji",
		Items:     emojiChoice,
		Templates: templates,
		Searcher:  searcher,
		Size:      4,
	}

	emojiNo, _, err := selectEmoji.Run()

	clipboard.WriteAll(el.Groups[groupNo].SubGroups[subGNo].Emojis[emojiNo].Symbol)

	fmt.Printf("You choose %q\n", el.Groups[groupNo].SubGroups[subGNo].Emojis[emojiNo].Symbol+" - it has now been copied to clipboard")
}
