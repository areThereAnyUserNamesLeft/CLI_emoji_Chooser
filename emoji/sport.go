package emoji

// MakeSportGroup returns the FaceGroup
func MakeSportGroup() Group {
	return Group{
		Name:       "sport",
		GroupNames: []string{"sport"},
		SubGroups: []subGroup{
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"fencing"},
				Emojis: []emoji{
					emoji{"🤺", []string{"person fencing"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"horse-racing"},
				Emojis: []emoji{
					emoji{"🏇", []string{"horse racing"}},
					emoji{"🏇🏻", []string{"horse racing: light skin tone"}},
					emoji{"🏇🏼", []string{"horse racing: medium-light skin tone"}},
					emoji{"🏇🏽", []string{"horse racing: medium skin tone"}},
					emoji{"🏇🏾", []string{"horse racing: medium-dark skin tone"}},
					emoji{"🏇🏿", []string{"horse racing: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"skier"},
				Emojis: []emoji{
					emoji{"⛷️", []string{"skier"}},
					emoji{"⛷", []string{"skier"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"snowboarder"},
				Emojis: []emoji{
					emoji{"🏂", []string{"snowboarder"}},
					emoji{"🏂🏻", []string{"snowboarder: light skin tone"}},
					emoji{"🏂🏼", []string{"snowboarder: medium-light skin tone"}},
					emoji{"🏂🏽", []string{"snowboarder: medium skin tone"}},
					emoji{"🏂🏾", []string{"snowboarder: medium-dark skin tone"}},
					emoji{"🏂🏿", []string{"snowboarder: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"golfing"},
				Emojis: []emoji{
					emoji{"🏌️", []string{"person golfing"}},
					emoji{"🏌", []string{"person golfing"}},
					emoji{"🏌🏻", []string{"person golfing: light skin tone"}},
					emoji{"🏌🏼", []string{"person golfing: medium-light skin tone"}},
					emoji{"🏌🏽", []string{"person golfing: medium skin tone"}},
					emoji{"🏌🏾", []string{"person golfing: medium-dark skin tone"}},
					emoji{"🏌🏿", []string{"person golfing: dark skin tone"}},
					emoji{"🏌️‍♂️", []string{"man golfing"}},
					emoji{"🏌‍♂️", []string{"man golfing"}},
					emoji{"🏌️‍♂", []string{"man golfing"}},
					emoji{"🏌‍♂", []string{"man golfing"}},
					emoji{"🏌🏻‍♂️", []string{"man golfing: light skin tone"}},
					emoji{"🏌🏻‍♂", []string{"man golfing: light skin tone"}},
					emoji{"🏌🏼‍♂️", []string{"man golfing: medium-light skin tone"}},
					emoji{"🏌🏼‍♂", []string{"man golfing: medium-light skin tone"}},
					emoji{"🏌🏽‍♂️", []string{"man golfing: medium skin tone"}},
					emoji{"🏌🏽‍♂", []string{"man golfing: medium skin tone"}},
					emoji{"🏌🏾‍♂️", []string{"man golfing: medium-dark skin tone"}},
					emoji{"🏌🏾‍♂", []string{"man golfing: medium-dark skin tone"}},
					emoji{"🏌🏿‍♂️", []string{"man golfing: dark skin tone"}},
					emoji{"🏌🏿‍♂", []string{"man golfing: dark skin tone"}},
					emoji{"🏌️‍♀️", []string{"woman golfing"}},
					emoji{"🏌‍♀️", []string{"woman golfing"}},
					emoji{"🏌️‍♀", []string{"woman golfing"}},
					emoji{"🏌‍♀", []string{"woman golfing"}},
					emoji{"🏌🏻‍♀️", []string{"woman golfing: light skin tone"}},
					emoji{"🏌🏻‍♀", []string{"woman golfing: light skin tone"}},
					emoji{"🏌🏼‍♀️", []string{"woman golfing: medium-light skin tone"}},
					emoji{"🏌🏼‍♀", []string{"woman golfing: medium-light skin tone"}},
					emoji{"🏌🏽‍♀️", []string{"woman golfing: medium skin tone"}},
					emoji{"🏌🏽‍♀", []string{"woman golfing: medium skin tone"}},
					emoji{"🏌🏾‍♀️", []string{"woman golfing: medium-dark skin tone"}},
					emoji{"🏌🏾‍♀", []string{"woman golfing: medium-dark skin tone"}},
					emoji{"🏌🏿‍♀️", []string{"woman golfing: dark skin tone"}},
					emoji{"🏌🏿‍♀", []string{"woman golfing: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"surfing"},
				Emojis: []emoji{
					emoji{"🏄", []string{"person surfing"}},
					emoji{"🏄🏻", []string{"person surfing: light skin tone"}},
					emoji{"🏄🏼", []string{"person surfing: medium-light skin tone"}},
					emoji{"🏄🏽", []string{"person surfing: medium skin tone"}},
					emoji{"🏄🏾", []string{"person surfing: medium-dark skin tone"}},
					emoji{"🏄🏿", []string{"person surfing: dark skin tone"}},
					emoji{"🏄‍♂️", []string{"man surfing"}},
					emoji{"🏄‍♂", []string{"man surfing"}},
					emoji{"🏄🏻‍♂️", []string{"man surfing: light skin tone"}},
					emoji{"🏄🏻‍♂", []string{"man surfing: light skin tone"}},
					emoji{"🏄🏼‍♂️", []string{"man surfing: medium-light skin tone"}},
					emoji{"🏄🏼‍♂", []string{"man surfing: medium-light skin tone"}},
					emoji{"🏄🏽‍♂️", []string{"man surfing: medium skin tone"}},
					emoji{"🏄🏽‍♂", []string{"man surfing: medium skin tone"}},
					emoji{"🏄🏾‍♂️", []string{"man surfing: medium-dark skin tone"}},
					emoji{"🏄🏾‍♂", []string{"man surfing: medium-dark skin tone"}},
					emoji{"🏄🏿‍♂️", []string{"man surfing: dark skin tone"}},
					emoji{"🏄🏿‍♂", []string{"man surfing: dark skin tone"}},
					emoji{"🏄‍♀️", []string{"woman surfing"}},
					emoji{"🏄‍♀", []string{"woman surfing"}},
					emoji{"🏄🏻‍♀️", []string{"woman surfing: light skin tone"}},
					emoji{"🏄🏻‍♀", []string{"woman surfing: light skin tone"}},
					emoji{"🏄🏼‍♀️", []string{"woman surfing: medium-light skin tone"}},
					emoji{"🏄🏼‍♀", []string{"woman surfing: medium-light skin tone"}},
					emoji{"🏄🏽‍♀️", []string{"woman surfing: medium skin tone"}},
					emoji{"🏄🏽‍♀", []string{"woman surfing: medium skin tone"}},
					emoji{"🏄🏾‍♀️", []string{"woman surfing: medium-dark skin tone"}},
					emoji{"🏄🏾‍♀", []string{"woman surfing: medium-dark skin tone"}},
					emoji{"🏄🏿‍♀️", []string{"woman surfing: dark skin tone"}},
					emoji{"🏄🏿‍♀", []string{"woman surfing: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"rowing"},
				Emojis: []emoji{
					emoji{"🚣", []string{"person rowing boat"}},
					emoji{"🚣🏻", []string{"person rowing boat: light skin tone"}},
					emoji{"🚣🏼", []string{"person rowing boat: medium-light skin tone"}},
					emoji{"🚣🏽", []string{"person rowing boat: medium skin tone"}},
					emoji{"🚣🏾", []string{"person rowing boat: medium-dark skin tone"}},
					emoji{"🚣🏿", []string{"person rowing boat: dark skin tone"}},
					emoji{"🚣‍♂️", []string{"man rowing boat"}},
					emoji{"🚣‍♂", []string{"man rowing boat"}},
					emoji{"🚣🏻‍♂️", []string{"man rowing boat: light skin tone"}},
					emoji{"🚣🏻‍♂", []string{"man rowing boat: light skin tone"}},
					emoji{"🚣🏼‍♂️", []string{"man rowing boat: medium-light skin tone"}},
					emoji{"🚣🏼‍♂", []string{"man rowing boat: medium-light skin tone"}},
					emoji{"🚣🏽‍♂️", []string{"man rowing boat: medium skin tone"}},
					emoji{"🚣🏽‍♂", []string{"man rowing boat: medium skin tone"}},
					emoji{"🚣🏾‍♂️", []string{"man rowing boat: medium-dark skin tone"}},
					emoji{"🚣🏾‍♂", []string{"man rowing boat: medium-dark skin tone"}},
					emoji{"🚣🏿‍♂️", []string{"man rowing boat: dark skin tone"}},
					emoji{"🚣🏿‍♂", []string{"man rowing boat: dark skin tone"}},
					emoji{"🚣‍♀️", []string{"woman rowing boat"}},
					emoji{"🚣‍♀", []string{"woman rowing boat"}},
					emoji{"🚣🏻‍♀️", []string{"woman rowing boat: light skin tone"}},
					emoji{"🚣🏻‍♀", []string{"woman rowing boat: light skin tone"}},
					emoji{"🚣🏼‍♀️", []string{"woman rowing boat: medium-light skin tone"}},
					emoji{"🚣🏼‍♀", []string{"woman rowing boat: medium-light skin tone"}},
					emoji{"🚣🏽‍♀️", []string{"woman rowing boat: medium skin tone"}},
					emoji{"🚣🏽‍♀", []string{"woman rowing boat: medium skin tone"}},
					emoji{"🚣🏾‍♀️", []string{"woman rowing boat: medium-dark skin tone"}},
					emoji{"🚣🏾‍♀", []string{"woman rowing boat: medium-dark skin tone"}},
					emoji{"🚣🏿‍♀️", []string{"woman rowing boat: dark skin tone"}},
					emoji{"🚣🏿‍♀", []string{"woman rowing boat: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"swimming"},
				Emojis: []emoji{
					emoji{"🏊", []string{"person swimming"}},
					emoji{"🏊🏻", []string{"person swimming: light skin tone"}},
					emoji{"🏊🏼", []string{"person swimming: medium-light skin tone"}},
					emoji{"🏊🏽", []string{"person swimming: medium skin tone"}},
					emoji{"🏊🏾", []string{"person swimming: medium-dark skin tone"}},
					emoji{"🏊🏿", []string{"person swimming: dark skin tone"}},
					emoji{"🏊‍♂️", []string{"man swimming"}},
					emoji{"🏊‍♂", []string{"man swimming"}},
					emoji{"🏊🏻‍♂️", []string{"man swimming: light skin tone"}},
					emoji{"🏊🏻‍♂", []string{"man swimming: light skin tone"}},
					emoji{"🏊🏼‍♂️", []string{"man swimming: medium-light skin tone"}},
					emoji{"🏊🏼‍♂", []string{"man swimming: medium-light skin tone"}},
					emoji{"🏊🏽‍♂️", []string{"man swimming: medium skin tone"}},
					emoji{"🏊🏽‍♂", []string{"man swimming: medium skin tone"}},
					emoji{"🏊🏾‍♂️", []string{"man swimming: medium-dark skin tone"}},
					emoji{"🏊🏾‍♂", []string{"man swimming: medium-dark skin tone"}},
					emoji{"🏊🏿‍♂️", []string{"man swimming: dark skin tone"}},
					emoji{"🏊🏿‍♂", []string{"man swimming: dark skin tone"}},
					emoji{"🏊‍♀️", []string{"woman swimming"}},
					emoji{"🏊‍♀", []string{"woman swimming"}},
					emoji{"🏊🏻‍♀️", []string{"woman swimming: light skin tone"}},
					emoji{"🏊🏻‍♀", []string{"woman swimming: light skin tone"}},
					emoji{"🏊🏼‍♀️", []string{"woman swimming: medium-light skin tone"}},
					emoji{"🏊🏼‍♀", []string{"woman swimming: medium-light skin tone"}},
					emoji{"🏊🏽‍♀️", []string{"woman swimming: medium skin tone"}},
					emoji{"🏊🏽‍♀", []string{"woman swimming: medium skin tone"}},
					emoji{"🏊🏾‍♀️", []string{"woman swimming: medium-dark skin tone"}},
					emoji{"🏊🏾‍♀", []string{"woman swimming: medium-dark skin tone"}},
					emoji{"🏊🏿‍♀️", []string{"woman swimming: dark skin tone"}},
					emoji{"🏊🏿‍♀", []string{"woman swimming: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"basketball"},
				Emojis: []emoji{
					emoji{"⛹️", []string{"person bouncing ball"}},
					emoji{"⛹", []string{"person bouncing ball"}},
					emoji{"⛹🏻", []string{"person bouncing ball: light skin tone"}},
					emoji{"⛹🏼", []string{"person bouncing ball: medium-light skin tone"}},
					emoji{"⛹🏽", []string{"person bouncing ball: medium skin tone"}},
					emoji{"⛹🏾", []string{"person bouncing ball: medium-dark skin tone"}},
					emoji{"⛹🏿", []string{"person bouncing ball: dark skin tone"}},
					emoji{"⛹️‍♂️", []string{"man bouncing ball"}},
					emoji{"⛹‍♂️", []string{"man bouncing ball"}},
					emoji{"⛹️‍♂", []string{"man bouncing ball"}},
					emoji{"⛹‍♂", []string{"man bouncing ball"}},
					emoji{"⛹🏻‍♂️", []string{"man bouncing ball: light skin tone"}},
					emoji{"⛹🏻‍♂", []string{"man bouncing ball: light skin tone"}},
					emoji{"⛹🏼‍♂️", []string{"man bouncing ball: medium-light skin tone"}},
					emoji{"⛹🏼‍♂", []string{"man bouncing ball: medium-light skin tone"}},
					emoji{"⛹🏽‍♂️", []string{"man bouncing ball: medium skin tone"}},
					emoji{"⛹🏽‍♂", []string{"man bouncing ball: medium skin tone"}},
					emoji{"⛹🏾‍♂️", []string{"man bouncing ball: medium-dark skin tone"}},
					emoji{"⛹🏾‍♂", []string{"man bouncing ball: medium-dark skin tone"}},
					emoji{"⛹🏿‍♂️", []string{"man bouncing ball: dark skin tone"}},
					emoji{"⛹🏿‍♂", []string{"man bouncing ball: dark skin tone"}},
					emoji{"⛹️‍♀️", []string{"woman bouncing ball"}},
					emoji{"⛹‍♀️", []string{"woman bouncing ball"}},
					emoji{"⛹️‍♀", []string{"woman bouncing ball"}},
					emoji{"⛹‍♀", []string{"woman bouncing ball"}},
					emoji{"⛹🏻‍♀️", []string{"woman bouncing ball: light skin tone"}},
					emoji{"⛹🏻‍♀", []string{"woman bouncing ball: light skin tone"}},
					emoji{"⛹🏼‍♀️", []string{"woman bouncing ball: medium-light skin tone"}},
					emoji{"⛹🏼‍♀", []string{"woman bouncing ball: medium-light skin tone"}},
					emoji{"⛹🏽‍♀️", []string{"woman bouncing ball: medium skin tone"}},
					emoji{"⛹🏽‍♀", []string{"woman bouncing ball: medium skin tone"}},
					emoji{"⛹🏾‍♀️", []string{"woman bouncing ball: medium-dark skin tone"}},
					emoji{"⛹🏾‍♀", []string{"woman bouncing ball: medium-dark skin tone"}},
					emoji{"⛹🏿‍♀️", []string{"woman bouncing ball: dark skin tone"}},
					emoji{"⛹🏿‍♀", []string{"woman bouncing ball: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"weights"},
				Emojis: []emoji{
					emoji{"🏋️", []string{"person lifting weights"}},
					emoji{"🏋", []string{"person lifting weights"}},
					emoji{"🏋🏻", []string{"person lifting weights: light skin tone"}},
					emoji{"🏋🏼", []string{"person lifting weights: medium-light skin tone"}},
					emoji{"🏋🏽", []string{"person lifting weights: medium skin tone"}},
					emoji{"🏋🏾", []string{"person lifting weights: medium-dark skin tone"}},
					emoji{"🏋🏿", []string{"person lifting weights: dark skin tone"}},
					emoji{"🏋️‍♂️", []string{"man lifting weights"}},
					emoji{"🏋‍♂️", []string{"man lifting weights"}},
					emoji{"🏋️‍♂", []string{"man lifting weights"}},
					emoji{"🏋‍♂", []string{"man lifting weights"}},
					emoji{"🏋🏻‍♂️", []string{"man lifting weights: light skin tone"}},
					emoji{"🏋🏻‍♂", []string{"man lifting weights: light skin tone"}},
					emoji{"🏋🏼‍♂️", []string{"man lifting weights: medium-light skin tone"}},
					emoji{"🏋🏼‍♂", []string{"man lifting weights: medium-light skin tone"}},
					emoji{"🏋🏽‍♂️", []string{"man lifting weights: medium skin tone"}},
					emoji{"🏋🏽‍♂", []string{"man lifting weights: medium skin tone"}},
					emoji{"🏋🏾‍♂️", []string{"man lifting weights: medium-dark skin tone"}},
					emoji{"🏋🏾‍♂", []string{"man lifting weights: medium-dark skin tone"}},
					emoji{"🏋🏿‍♂️", []string{"man lifting weights: dark skin tone"}},
					emoji{"🏋🏿‍♂", []string{"man lifting weights: dark skin tone"}},
					emoji{"🏋️‍♀️", []string{"woman lifting weights"}},
					emoji{"🏋‍♀️", []string{"woman lifting weights"}},
					emoji{"🏋️‍♀", []string{"woman lifting weights"}},
					emoji{"🏋‍♀", []string{"woman lifting weights"}},
					emoji{"🏋🏻‍♀️", []string{"woman lifting weights: light skin tone"}},
					emoji{"🏋🏻‍♀", []string{"woman lifting weights: light skin tone"}},
					emoji{"🏋🏼‍♀️", []string{"woman lifting weights: medium-light skin tone"}},
					emoji{"🏋🏼‍♀", []string{"woman lifting weights: medium-light skin tone"}},
					emoji{"🏋🏽‍♀️", []string{"woman lifting weights: medium skin tone"}},
					emoji{"🏋🏽‍♀", []string{"woman lifting weights: medium skin tone"}},
					emoji{"🏋🏾‍♀️", []string{"woman lifting weights: medium-dark skin tone"}},
					emoji{"🏋🏾‍♀", []string{"woman lifting weights: medium-dark skin tone"}},
					emoji{"🏋🏿‍♀️", []string{"woman lifting weights: dark skin tone"}},
					emoji{"🏋🏿‍♀", []string{"woman lifting weights: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"biking"},
				Emojis: []emoji{
					emoji{"🚴", []string{"person biking"}},
					emoji{"🚴🏻", []string{"person biking: light skin tone"}},
					emoji{"🚴🏼", []string{"person biking: medium-light skin tone"}},
					emoji{"🚴🏽", []string{"person biking: medium skin tone"}},
					emoji{"🚴🏾", []string{"person biking: medium-dark skin tone"}},
					emoji{"🚴🏿", []string{"person biking: dark skin tone"}},
					emoji{"🚴‍♂️", []string{"man biking"}},
					emoji{"🚴‍♂", []string{"man biking"}},
					emoji{"🚴🏻‍♂️", []string{"man biking: light skin tone"}},
					emoji{"🚴🏻‍♂", []string{"man biking: light skin tone"}},
					emoji{"🚴🏼‍♂️", []string{"man biking: medium-light skin tone"}},
					emoji{"🚴🏼‍♂", []string{"man biking: medium-light skin tone"}},
					emoji{"🚴🏽‍♂️", []string{"man biking: medium skin tone"}},
					emoji{"🚴🏽‍♂", []string{"man biking: medium skin tone"}},
					emoji{"🚴🏾‍♂️", []string{"man biking: medium-dark skin tone"}},
					emoji{"🚴🏾‍♂", []string{"man biking: medium-dark skin tone"}},
					emoji{"🚴🏿‍♂️", []string{"man biking: dark skin tone"}},
					emoji{"🚴🏿‍♂", []string{"man biking: dark skin tone"}},
					emoji{"🚴‍♀️", []string{"woman biking"}},
					emoji{"🚴‍♀", []string{"woman biking"}},
					emoji{"🚴🏻‍♀️", []string{"woman biking: light skin tone"}},
					emoji{"🚴🏻‍♀", []string{"woman biking: light skin tone"}},
					emoji{"🚴🏼‍♀️", []string{"woman biking: medium-light skin tone"}},
					emoji{"🚴🏼‍♀", []string{"woman biking: medium-light skin tone"}},
					emoji{"🚴🏽‍♀️", []string{"woman biking: medium skin tone"}},
					emoji{"🚴🏽‍♀", []string{"woman biking: medium skin tone"}},
					emoji{"🚴🏾‍♀️", []string{"woman biking: medium-dark skin tone"}},
					emoji{"🚴🏾‍♀", []string{"woman biking: medium-dark skin tone"}},
					emoji{"🚴🏿‍♀️", []string{"woman biking: dark skin tone"}},
					emoji{"🚴🏿‍♀", []string{"woman biking: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"mountain-biking"},
				Emojis: []emoji{
					emoji{"🚵", []string{"person mountain biking"}},
					emoji{"🚵🏻", []string{"person mountain biking: light skin tone"}},
					emoji{"🚵🏼", []string{"person mountain biking: medium-light skin tone"}},
					emoji{"🚵🏽", []string{"person mountain biking: medium skin tone"}},
					emoji{"🚵🏾", []string{"person mountain biking: medium-dark skin tone"}},
					emoji{"🚵🏿", []string{"person mountain biking: dark skin tone"}},
					emoji{"🚵‍♂️", []string{"man mountain biking"}},
					emoji{"🚵‍♂", []string{"man mountain biking"}},
					emoji{"🚵🏻‍♂️", []string{"man mountain biking: light skin tone"}},
					emoji{"🚵🏻‍♂", []string{"man mountain biking: light skin tone"}},
					emoji{"🚵🏼‍♂️", []string{"man mountain biking: medium-light skin tone"}},
					emoji{"🚵🏼‍♂", []string{"man mountain biking: medium-light skin tone"}},
					emoji{"🚵🏽‍♂️", []string{"man mountain biking: medium skin tone"}},
					emoji{"🚵🏽‍♂", []string{"man mountain biking: medium skin tone"}},
					emoji{"🚵🏾‍♂️", []string{"man mountain biking: medium-dark skin tone"}},
					emoji{"🚵🏾‍♂", []string{"man mountain biking: medium-dark skin tone"}},
					emoji{"🚵🏿‍♂️", []string{"man mountain biking: dark skin tone"}},
					emoji{"🚵🏿‍♂", []string{"man mountain biking: dark skin tone"}},
					emoji{"🚵‍♀️", []string{"woman mountain biking"}},
					emoji{"🚵‍♀", []string{"woman mountain biking"}},
					emoji{"🚵🏻‍♀️", []string{"woman mountain biking: light skin tone"}},
					emoji{"🚵🏻‍♀", []string{"woman mountain biking: light skin tone"}},
					emoji{"🚵🏼‍♀️", []string{"woman mountain biking: medium-light skin tone"}},
					emoji{"🚵🏼‍♀", []string{"woman mountain biking: medium-light skin tone"}},
					emoji{"🚵🏽‍♀️", []string{"woman mountain biking: medium skin tone"}},
					emoji{"🚵🏽‍♀", []string{"woman mountain biking: medium skin tone"}},
					emoji{"🚵🏾‍♀️", []string{"woman mountain biking: medium-dark skin tone"}},
					emoji{"🚵🏾‍♀", []string{"woman mountain biking: medium-dark skin tone"}},
					emoji{"🚵🏿‍♀️", []string{"woman mountain biking: dark skin tone"}},
					emoji{"🚵🏿‍♀", []string{"woman mountain biking: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"racing-car"},
				Emojis: []emoji{
					emoji{"🏎️", []string{"racing car"}},
					emoji{"🏎", []string{"racing car"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"motorcycle"},
				Emojis: []emoji{
					emoji{"🏍️", []string{"motorcycle"}},
					emoji{"🏍", []string{"motorcycle"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"cartwheeling"},
				Emojis: []emoji{
					emoji{"🤸", []string{"person cartwheeling"}},
					emoji{"🤸🏻", []string{"person cartwheeling: light skin tone"}},
					emoji{"🤸🏼", []string{"person cartwheeling: medium-light skin tone"}},
					emoji{"🤸🏽", []string{"person cartwheeling: medium skin tone"}},
					emoji{"🤸🏾", []string{"person cartwheeling: medium-dark skin tone"}},
					emoji{"🤸🏿", []string{"person cartwheeling: dark skin tone"}},
					emoji{"🤸‍♂️", []string{"man cartwheeling"}},
					emoji{"🤸‍♂", []string{"man cartwheeling"}},
					emoji{"🤸🏻‍♂️", []string{"man cartwheeling: light skin tone"}},
					emoji{"🤸🏻‍♂", []string{"man cartwheeling: light skin tone"}},
					emoji{"🤸🏼‍♂️", []string{"man cartwheeling: medium-light skin tone"}},
					emoji{"🤸🏼‍♂", []string{"man cartwheeling: medium-light skin tone"}},
					emoji{"🤸🏽‍♂️", []string{"man cartwheeling: medium skin tone"}},
					emoji{"🤸🏽‍♂", []string{"man cartwheeling: medium skin tone"}},
					emoji{"🤸🏾‍♂️", []string{"man cartwheeling: medium-dark skin tone"}},
					emoji{"🤸🏾‍♂", []string{"man cartwheeling: medium-dark skin tone"}},
					emoji{"🤸🏿‍♂️", []string{"man cartwheeling: dark skin tone"}},
					emoji{"🤸🏿‍♂", []string{"man cartwheeling: dark skin tone"}},
					emoji{"🤸‍♀️", []string{"woman cartwheeling"}},
					emoji{"🤸‍♀", []string{"woman cartwheeling"}},
					emoji{"🤸🏻‍♀️", []string{"woman cartwheeling: light skin tone"}},
					emoji{"🤸🏻‍♀", []string{"woman cartwheeling: light skin tone"}},
					emoji{"🤸🏼‍♀️", []string{"woman cartwheeling: medium-light skin tone"}},
					emoji{"🤸🏼‍♀", []string{"woman cartwheeling: medium-light skin tone"}},
					emoji{"🤸🏽‍♀️", []string{"woman cartwheeling: medium skin tone"}},
					emoji{"🤸🏽‍♀", []string{"woman cartwheeling: medium skin tone"}},
					emoji{"🤸🏾‍♀️", []string{"woman cartwheeling: medium-dark skin tone"}},
					emoji{"🤸🏾‍♀", []string{"woman cartwheeling: medium-dark skin tone"}},
					emoji{"🤸🏿‍♀️", []string{"woman cartwheeling: dark skin tone"}},
					emoji{"🤸🏿‍♀", []string{"woman cartwheeling: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"wrestling"},
				Emojis: []emoji{
					emoji{"🤼", []string{"people wrestling"}},
					emoji{"🤼‍♂️", []string{"men wrestling"}},
					emoji{"🤼‍♂", []string{"men wrestling"}},
					emoji{"🤼‍♀️", []string{"women wrestling"}},
					emoji{"🤼‍♀", []string{"women wrestling"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"water-polo"},
				Emojis: []emoji{
					emoji{"🤽", []string{"person playing water polo"}},
					emoji{"🤽🏻", []string{"person playing water polo: light skin tone"}},
					emoji{"🤽🏼", []string{"person playing water polo: medium-light skin tone"}},
					emoji{"🤽🏽", []string{"person playing water polo: medium skin tone"}},
					emoji{"🤽🏾", []string{"person playing water polo: medium-dark skin tone"}},
					emoji{"🤽🏿", []string{"person playing water polo: dark skin tone"}},
					emoji{"🤽‍♂️", []string{"man playing water polo"}},
					emoji{"🤽‍♂", []string{"man playing water polo"}},
					emoji{"🤽🏻‍♂️", []string{"man playing water polo: light skin tone"}},
					emoji{"🤽🏻‍♂", []string{"man playing water polo: light skin tone"}},
					emoji{"🤽🏼‍♂️", []string{"man playing water polo: medium-light skin tone"}},
					emoji{"🤽🏼‍♂", []string{"man playing water polo: medium-light skin tone"}},
					emoji{"🤽🏽‍♂️", []string{"man playing water polo: medium skin tone"}},
					emoji{"🤽🏽‍♂", []string{"man playing water polo: medium skin tone"}},
					emoji{"🤽🏾‍♂️", []string{"man playing water polo: medium-dark skin tone"}},
					emoji{"🤽🏾‍♂", []string{"man playing water polo: medium-dark skin tone"}},
					emoji{"🤽🏿‍♂️", []string{"man playing water polo: dark skin tone"}},
					emoji{"🤽🏿‍♂", []string{"man playing water polo: dark skin tone"}},
					emoji{"🤽‍♀️", []string{"woman playing water polo"}},
					emoji{"🤽‍♀", []string{"woman playing water polo"}},
					emoji{"🤽🏻‍♀️", []string{"woman playing water polo: light skin tone"}},
					emoji{"🤽🏻‍♀", []string{"woman playing water polo: light skin tone"}},
					emoji{"🤽🏼‍♀️", []string{"woman playing water polo: medium-light skin tone"}},
					emoji{"🤽🏼‍♀", []string{"woman playing water polo: medium-light skin tone"}},
					emoji{"🤽🏽‍♀️", []string{"woman playing water polo: medium skin tone"}},
					emoji{"🤽🏽‍♀", []string{"woman playing water polo: medium skin tone"}},
					emoji{"🤽🏾‍♀️", []string{"woman playing water polo: medium-dark skin tone"}},
					emoji{"🤽🏾‍♀", []string{"woman playing water polo: medium-dark skin tone"}},
					emoji{"🤽🏿‍♀️", []string{"woman playing water polo: dark skin tone"}},
					emoji{"🤽🏿‍♀", []string{"woman playing water polo: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"handball"},
				Emojis: []emoji{
					emoji{"🤾", []string{"person playing handball"}},
					emoji{"🤾🏻", []string{"person playing handball: light skin tone"}},
					emoji{"🤾🏼", []string{"person playing handball: medium-light skin tone"}},
					emoji{"🤾🏽", []string{"person playing handball: medium skin tone"}},
					emoji{"🤾🏾", []string{"person playing handball: medium-dark skin tone"}},
					emoji{"🤾🏿", []string{"person playing handball: dark skin tone"}},
					emoji{"🤾‍♂️", []string{"man playing handball"}},
					emoji{"🤾‍♂", []string{"man playing handball"}},
					emoji{"🤾🏻‍♂️", []string{"man playing handball: light skin tone"}},
					emoji{"🤾🏻‍♂", []string{"man playing handball: light skin tone"}},
					emoji{"🤾🏼‍♂️", []string{"man playing handball: medium-light skin tone"}},
					emoji{"🤾🏼‍♂", []string{"man playing handball: medium-light skin tone"}},
					emoji{"🤾🏽‍♂️", []string{"man playing handball: medium skin tone"}},
					emoji{"🤾🏽‍♂", []string{"man playing handball: medium skin tone"}},
					emoji{"🤾🏾‍♂️", []string{"man playing handball: medium-dark skin tone"}},
					emoji{"🤾🏾‍♂", []string{"man playing handball: medium-dark skin tone"}},
					emoji{"🤾🏿‍♂️", []string{"man playing handball: dark skin tone"}},
					emoji{"🤾🏿‍♂", []string{"man playing handball: dark skin tone"}},
					emoji{"🤾‍♀️", []string{"woman playing handball"}},
					emoji{"🤾‍♀", []string{"woman playing handball"}},
					emoji{"🤾🏻‍♀️", []string{"woman playing handball: light skin tone"}},
					emoji{"🤾🏻‍♀", []string{"woman playing handball: light skin tone"}},
					emoji{"🤾🏼‍♀️", []string{"woman playing handball: medium-light skin tone"}},
					emoji{"🤾🏼‍♀", []string{"woman playing handball: medium-light skin tone"}},
					emoji{"🤾🏽‍♀️", []string{"woman playing handball: medium skin tone"}},
					emoji{"🤾🏽‍♀", []string{"woman playing handball: medium skin tone"}},
					emoji{"🤾🏾‍♀️", []string{"woman playing handball: medium-dark skin tone"}},
					emoji{"🤾🏾‍♀", []string{"woman playing handball: medium-dark skin tone"}},
					emoji{"🤾🏿‍♀️", []string{"woman playing handball: dark skin tone"}},
					emoji{"🤾🏿‍♀", []string{"woman playing handball: dark skin tone"}},
				},
			},
			subGroup{
				Group:         "sport",
				SubGroupNames: []string{"juggling"},
				Emojis: []emoji{
					emoji{"🤹", []string{"person juggling"}},
					emoji{"🤹🏻", []string{"person juggling: light skin tone"}},
					emoji{"🤹🏼", []string{"person juggling: medium-light skin tone"}},
					emoji{"🤹🏽", []string{"person juggling: medium skin tone"}},
					emoji{"🤹🏾", []string{"person juggling: medium-dark skin tone"}},
					emoji{"🤹🏿", []string{"person juggling: dark skin tone"}},
					emoji{"🤹‍♂️", []string{"man juggling"}},
					emoji{"🤹‍♂", []string{"man juggling"}},
					emoji{"🤹🏻‍♂️", []string{"man juggling: light skin tone"}},
					emoji{"🤹🏻‍♂", []string{"man juggling: light skin tone"}},
					emoji{"🤹🏼‍♂️", []string{"man juggling: medium-light skin tone"}},
					emoji{"🤹🏼‍♂", []string{"man juggling: medium-light skin tone"}},
					emoji{"🤹🏽‍♂️", []string{"man juggling: medium skin tone"}},
					emoji{"🤹🏽‍♂", []string{"man juggling: medium skin tone"}},
					emoji{"🤹🏾‍♂️", []string{"man juggling: medium-dark skin tone"}},
					emoji{"🤹🏾‍♂", []string{"man juggling: medium-dark skin tone"}},
					emoji{"🤹🏿‍♂️", []string{"man juggling: dark skin tone"}},
					emoji{"🤹🏿‍♂", []string{"man juggling: dark skin tone"}},
					emoji{"🤹‍♀️", []string{"woman juggling"}},
					emoji{"🤹‍♀", []string{"woman juggling"}},
					emoji{"🤹🏻‍♀️", []string{"woman juggling: light skin tone"}},
					emoji{"🤹🏻‍♀", []string{"woman juggling: light skin tone"}},
					emoji{"🤹🏼‍♀️", []string{"woman juggling: medium-light skin tone"}},
					emoji{"🤹🏼‍♀", []string{"woman juggling: medium-light skin tone"}},
					emoji{"🤹🏽‍♀️", []string{"woman juggling: medium skin tone"}},
					emoji{"🤹🏽‍♀", []string{"woman juggling: medium skin tone"}},
					emoji{"🤹🏾‍♀️", []string{"woman juggling: medium-dark skin tone"}},
					emoji{"🤹🏾‍♀", []string{"woman juggling: medium-dark skin tone"}},
					emoji{"🤹🏿‍♀️", []string{"woman juggling: dark skin tone"}},
					emoji{"🤹🏿‍♀", []string{"woman juggling: dark skin tone"}},
				},
			},
		},
	}
}
