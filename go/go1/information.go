package main

type character struct {
	name                                 string
	level, hp, atk, def, dex, exp, expup int
}

var item character = character{"裝備", 0, 0, 0, 0, 0, 0, 0}
var p0 character = character{"基本值", 1, 5, 2, 2, 2, 10, 0}
var p1 character = character{"勇者", 1, 5, 2, 2, 2, 0, 0}
var m1 character = character{"史萊姆", 1, 10, 1, 1, 1, 1, 0}
var m2 character = character{"蝙蝠", 1, 5, 2, 1, 5, 2, 0}
var m3 character = character{"哥布林", 1, 8, 3, 2, 2, 3, 0}
var m4 character = character{"巨鼠", 1, 9, 1, 4, 2, 4, 0}
var m5 character = character{"殭屍", 1, 15, 2, 2, 1, 5, 0}
var b1 character = character{"魔龍", 30, 1000, 100, 80, 50, 10000, 0}

func itemlevel(x int) string { ////////////道具等級資訊
	var y string
	switch {
	case x <= 5:
		y = "受損的"
	case x <= 7:
		y = "老舊的"
	case x <= 9:
		y = "普通的"
	case x <= 11:
		y = "稀有的"
	case x <= 13:
		y = "罕見的"
	case x <= 15:
		y = "史詩的"
	case x <= 17:
		y = "古代的"
	case x <= 19:
		y = "傳說的"
	case x >= 20:
		y = "神的"
	}
	return y
}

// func updatemtr(mtr character, level int, hp int, atk int, def int, dex int, exp int) { ////更新怪物資訊
// 	mtr.level = level
// 	mtr.hp = hp
// 	mtr.atk = atk
// 	mtr.def = def
// 	mtr.dex = dex
// 	mtr.exp = exp
// }

func updatep1() { //////////////更新戰鬥資訊
	p1.level = p0.level
	p1.hp = p0.hp
	p1.atk = p0.atk + item.atk
	p1.def = p0.def + item.def
	p1.dex = p0.dex + item.dex
}

func updatepitem() { //////更新裝備資訊
	p1.atk = p0.atk + item.atk
	p1.def = p0.def + item.def
	p1.dex = p0.dex + item.dex
}
