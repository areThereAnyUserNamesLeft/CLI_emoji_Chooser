package emoji

type EmojiList struct {
	Groups []Group
}

type Group struct {
	Name       string
	GroupNames []string
	SubGroups  []subGroup
}

type subGroup struct {
	Group         string
	SubGroupNames []string
	Emojis        []emoji
}

type emoji struct {
	Symbol string
	Names  []string
}

func (el *EmojiList) GetGroupNames() (gn []string) {
	for g := range el.Groups {
		gn = append(gn, el.Groups[g].Name)
	}
	return gn
}
