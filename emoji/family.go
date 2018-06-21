package emoji

// MakeFamilyGroup returns the familyGroup
func MakeFamilyGroup() Group {
	return Group{
		Name:       "family",
		GroupNames: []string{"family"},
		SubGroups: []subGroup{
			subGroup{
				Group:         "family",
				SubGroupNames: []string{"holding-hands"},
				Emojis: []emoji{
					emoji{"👫", []string{"man and woman holding hands"}},
					emoji{"👬", []string{"two men holding hands"}},
					emoji{"👭", []string{"two women holding hands"}},
				},
			},
			subGroup{
				Group:         "family",
				SubGroupNames: []string{"kiss"},
				Emojis: []emoji{
					emoji{"💏", []string{"kiss"}},
					emoji{"👩‍❤️‍💋‍👨", []string{"kiss: woman, man"}},
					emoji{"👩‍❤‍💋‍👨", []string{"kiss: woman, man"}},
					emoji{"👨‍❤️‍💋‍👨", []string{"kiss: man, man"}},
					emoji{"👨‍❤‍💋‍👨", []string{"kiss: man, man"}},
					emoji{"👩‍❤️‍💋‍👩", []string{"kiss: woman, woman"}},
					emoji{"👩‍❤‍💋‍👩", []string{"kiss: woman, woman"}},
				},
			},
			subGroup{
				Group:         "family",
				SubGroupNames: []string{"couple-with-heart"},
				Emojis: []emoji{
					emoji{"💑", []string{"couple with heart"}},
					emoji{"👩‍❤️‍👨", []string{"couple with heart: woman, man"}},
					emoji{"👩‍❤‍👨", []string{"couple with heart: woman, man"}},
					emoji{"👨‍❤️‍👨", []string{"couple with heart: man, man"}},
					emoji{"👨‍❤‍👨", []string{"couple with heart: man, man"}},
					emoji{"👩‍❤️‍👩", []string{"couple with heart: woman, woman"}},
					emoji{"👩‍❤‍👩", []string{"couple with heart: woman, woman"}},
				},
			},
			subGroup{
				Group:         "family",
				SubGroupNames: []string{"family group"},
				Emojis: []emoji{
					emoji{"👪", []string{"family"}},
					emoji{"👨‍👩‍👦", []string{"family: man, woman, boy"}},
					emoji{"👨‍👩‍👧", []string{"family: man, woman, girl"}},
					emoji{"👨‍👩‍👧‍👦", []string{"family: man, woman, girl, boy"}},
					emoji{"👨‍👩‍👦‍👦", []string{"family: man, woman, boy, boy"}},
					emoji{"👨‍👩‍👧‍👧", []string{"family: man, woman, girl, girl"}},
					emoji{"👨‍👨‍👦", []string{"family: man, man, boy"}},
					emoji{"👨‍👨‍👧", []string{"family: man, man, girl"}},
					emoji{"👨‍👨‍👧‍👦", []string{"family: man, man, girl, boy"}},
					emoji{"👨‍👨‍👦‍👦", []string{"family: man, man, boy, boy"}},
					emoji{"👨‍👨‍👧‍👧", []string{"family: man, man, girl, girl"}},
					emoji{"👩‍👩‍👦", []string{"family: woman, woman, boy"}},
					emoji{"👩‍👩‍👧", []string{"family: woman, woman, girl"}},
					emoji{"👩‍👩‍👧‍👦", []string{"family: woman, woman, girl, boy"}},
					emoji{"👩‍👩‍👦‍👦", []string{"family: woman, woman, boy, boy"}},
					emoji{"👩‍👩‍👧‍👧", []string{"family: woman, woman, girl, girl"}},
					emoji{"👨‍👦", []string{"family: man, boy"}},
					emoji{"👨‍👦‍👦", []string{"family: man, boy, boy"}},
					emoji{"👨‍👧", []string{"family: man, girl"}},
					emoji{"👨‍👧‍👦", []string{"family: man, girl, boy"}},
					emoji{"👨‍👧‍👧", []string{"family: man, girl, girl"}},
					emoji{"👩‍👦", []string{"family: woman, boy"}},
					emoji{"👩‍👦‍👦", []string{"family: woman, boy, boy"}},
					emoji{"👩‍👧", []string{"family: woman, girl"}},
					emoji{"👩‍👧‍👦", []string{"family: woman, girl, boy"}},
					emoji{"👩‍👧‍👧", []string{"family: woman, girl, girl"}},
				},
			},
		},
	}
}
