package main

import (
	f "fmt"
)

func resetall() { /////////////重置遊戲
	f.Println("要開始新的遊戲嗎?")
	f.Println("1.確定 2.先不要")
	switch userscan() {
	case "1":
		item = character{"裝備", 0, 0, 0, 0, 0, 0, 0}
		p0 = character{"基本值", 1, 5, 2, 2, 2, 10, 0}
		p1 = character{"勇者", 1, 5, 2, 2, 2, 0, 0}
		m1 = character{"史萊姆", 1, 10, 1, 1, 1, 1, 0}
		m2 = character{"蝙蝠", 1, 5, 2, 1, 5, 2, 0}
		m3 = character{"哥布林", 1, 8, 3, 2, 2, 3, 0}
		m4 = character{"巨鼠", 1, 9, 1, 4, 2, 4, 0}
		m5 = character{"殭屍", 1, 15, 2, 2, 1, 5, 0}
		b1 = character{"魔龍", 30, 1000, 100, 80, 50, 10000, 0}
	case "2":
		title()
	}
}

func chooseitem(level int) { //////////////////抽選道具
	x, y, z := roll3(level + 10)
	if b1.level <= 30 && (x > 25 || y > 25 || z > 25) { /////////////第一次打贏BOSS前裝備能力無法超過25
		for x > 25 || y > 25 || z > 25 {
			x, y, z = roll3(level + 10)
		}
	}
	f.Println("選吧~你要哪一件呢~") ///////////////////////////////////////////////////////////////////////////////裝備項目待增加
	f.Printf("1.%s大地之劍 2.%s森林之鎧 3.%s風之長靴 4.都不要\n", itemlevel(x), itemlevel(y), itemlevel(z))
	switch userscan() {
	case "1":
		item.atk = x
		getitem(itemlevel(x), "大地之劍")
	case "2":
		item.def = y
		getitem(itemlevel(y), "森林之鎧")
	case "3":
		item.dex = z
		getitem(itemlevel(z), "風之長靴")
	default:
		break
	}
	shop()
}

func questchoose(x int) string { ////////////任務選擇
	var qname string
	switch x {
	case 1:
		qname = m1.name
	case 2:
		qname = m2.name
	case 3:
		qname = m3.name
	case 4:
		qname = m4.name
	case 5:
		qname = m5.name
	}
	return qname
}

func fight() { /////////////////////////////進入戰鬥
	f.Println("你遭遇了:")
	var x int
	switch {
	case p1.level <= 30: ////30等前只會遇到小怪
		x = 5
	case p1.level > 30: /////30等候會遇到魔龍
		x = 7
	}
	y := roll(x)
	Mtr(y)
}

func qusetfight(x string) { ////////////////任務戰鬥
	y := roll(6)
	hpup := y
	f.Printf("請消滅%d隻%s!\n", y, x)
	for y != 0 {
		switch x {
		case m1.name:
			Mtr(1)
		case m2.name:
			Mtr(2)
		case m3.name:
			Mtr(3)
		case m4.name:
			Mtr(4)
		case m5.name:
			Mtr(5)
		}
		y--
	}
	if y == 0 {
		f.Println("完成狩獵任務。")
		f.Println("血量增加", hpup)
		p0.hp += hpup
	}
	groap()
}

func fighting(mtr character) { /////戰鬥過程
	f.Printf("%11s\n   等級: %d   血量: %d\n攻擊: %d 防禦: %d 敏捷: %d\n\n", mtr.name, mtr.level, mtr.hp, mtr.atk, mtr.def, mtr.dex)
	x := true
	for mtr.hp > 0 && p1.hp > 0 && x {
		pdef, patk := p1.def, p1.atk
		mdef, matk := mtr.def, mtr.atk
		f.Printf("勇者 等級:%d 血量:%d 攻擊:%d 防禦:%d 敏捷:%d\n", p1.level, p1.hp, patk, pdef, p1.dex)
		f.Println("1.攻擊 2.防禦 3.逃跑")
		switch userscan() {
		case "1":
			///////////////////////////////////////////////////////////////////////////////////////////////////////////////rouge技能能選項(待補)
			pd := roll(p1.dex + p1.level) ///////等級+敏捷當命中
			f.Printf("命中點數為%d，%s敏捷率為%d，你骰出了%d!\n", p0.level+p1.dex, mtr.name, mtr.dex, pd)
			if pd >= mtr.dex { ////命中
				if pd >= mtr.dex*3 { /////爆擊判定
					f.Println("判定爆擊!!進行", patk, "點攻擊!!")
					patk *= 2
				} else {
					f.Println("進行", patk, "點攻擊!")
				}
				if patk >= mtr.def {
					f.Printf("%s防禦為%d，造成%d點傷害!\n", mtr.name, mtr.def, patk-mtr.def)
					mtr.hp -= patk - mtr.def
				} else {
					f.Printf("傷害不足...，%s回復 1 點血量。\n", mtr.name)
					mtr.hp++
				}
				f.Println(mtr.name, "剩餘血量:", mtr.hp)
			} else { //////敏捷
				f.Println(mtr.name, "閃過了你的攻擊!")
			}
		case "2": /////////////防禦姿態可作為回復手段
			f.Println("勇者採取防禦姿勢!")
			pdef *= 2
		case "3":
			x = false
		default:
		}
		if mtr.hp > 0 && x {
			switch roll(2) {
			case 1:
				md := roll(mtr.dex + mtr.level) //////怪物命中
				f.Printf("你的敏捷點數為%d，%s骰出了%d!\n", p1.dex, mtr.name, md)
				if md >= p1.dex { //////命中
					if md >= p1.dex*2 { //////怪物爆擊判定
						f.Println(mtr.name, "判定爆擊!!進行", matk, "點攻擊!!")
						matk *= 2
					} else {
						f.Println(mtr.name, "對你進行", matk, "點攻擊!")
					}
					if matk >= p1.def { //////攻擊超過防禦進行傷害判定
						f.Printf("你的防禦為%d，造成%d點傷害!\n", p1.def, matk-p1.def)
						p1.hp -= matk - p1.def
					} else { ///攻擊低於防禦回復一點血量
						f.Println("傷害不足...，你回復 1 點血量。")
						p1.hp++
					}
				} else {
					f.Printf("你閃過了 %s 的攻擊!!\n", mtr.name)
				}
				f.Println("你剩餘血量:", p1.hp)
			default: //////////////怪物防禦也可作為回復手段
				f.Println(mtr.name, "採取防禦姿勢!")
				mdef *= 2
			}
		}
	}
	if mtr.hp <= 0 {
		p0.expup += mtr.exp
	}
	for p0.expup >= p0.exp { ///////////////////////升級計算
		levelup()
	}
	if p1.hp >= p0.hp { //////////////////////////戰鬥中超量的血回復正常
		p1.hp = p0.hp

	} else if p1.hp <= 0 { //////////////死亡
		dead()
	}
}

func Mtr(x int) { /////////遭遇
	switch {
	case x <= 1:
		fighting(m1)
		m1.expup += 1
		if m1.expup >= 3 {
			m1.level, m1.hp, m1.atk, m1.def, m1.dex, m1.exp = mtrlevelup(m1)
			m1.expup = 0
		}
	case x == 2:
		fighting(m2)
		m2.expup += 1
		if m2.expup >= 3 {
			m2.level, m2.hp, m2.atk, m2.def, m2.dex, m2.exp = mtrlevelup(m2)
			m2.expup = 0
		}
	case x == 3:
		fighting(m3)
		m3.expup += 1
		if m3.expup >= 3 {
			m3.level, m3.hp, m3.atk, m3.def, m3.dex, m3.exp = mtrlevelup(m3)
			m3.expup = 0
		}
	case x == 4:
		fighting(m4)
		m4.expup += 1
		if m4.expup >= 3 {
			m4.level, m4.hp, m4.atk, m4.def, m4.dex, m4.exp = mtrlevelup(m4)
			m4.expup = 0
		}
	case x == 5:
		fighting(m5)
		m5.expup += 1
		if m5.expup >= 3 {
			m5.level, m5.hp, m5.atk, m5.def, m5.dex, m5.exp = mtrlevelup(m5)
			m5.expup = 0
		}
	case x > 5:
		fighting(b1)
		if b1.expup >= 1 {
			b1.level, b1.hp, b1.atk, b1.def, b1.dex, b1.exp = bosslevelup(b1)
			b1.expup = 0
		}
		endgame()

	}
}
