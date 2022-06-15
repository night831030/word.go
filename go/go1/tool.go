package main

import (
	f "fmt"
	r "math/rand"
	t "time"
)

func roll(x int) int { //骰子
	r.Seed(t.Now().UnixNano())
	return r.Intn(100000000000000)%x + 1
}

func roll3(x int) (int, int, int) { ///////////////////三個亂數
	i, j, k := 0, 0, 0
	for i == j || i == k {
		i = roll(100000)
		for j == k {
			j = roll(100000)
			for k == 0 {
				k = roll(100000)
			}
		}
	}
	return ((i % x) + 1), ((j % x) + 1), ((k % x) + 1)
}

func userscan() string { //使用者輸入
	var user string
	f.Scan(&user)
	return user
	// var x string /////////////////////////測試用
	// switch roll(3) {
	// case 1:
	// 	x = "1"
	// case 2:
	// 	x = "2"
	// case 3:
	// 	x = "3"
	// default:
	// 	x = "0"
	// }
	// return x
}

func getitem(l string, n string) { ////////////道具獲得
	f.Print("獲得了 " + l + n + "。\n")
	updatepitem()
}

func levelup() { ////////////////////////角色升級
	f.Println("勇者升級了!!!")
	p0.level += 1
	p0.hp += 5
	f.Println("血量+5")
	f.Println("選擇能力:") //////////////////////////////////////////////////////////待補項目
	choose()
	p0.expup -= p0.exp
	p0.exp += p0.level * 2
	updatep1()
}

func choose() { ///////////////////////////////////////////////////////////////////////////////////////////////選擇能力(待修)
	var x, y, z int
	x, y, z = roll3(3)
	f.Print("1.")
	f.Print(showskill(x))
	f.Print(" 2.")
	f.Print(showskill(y))
	f.Print(" 3.")
	f.Print(showskill(z))
	f.Println("")
	switch userscan() {
	case "1":
		getskill(showskill(x))
	case "2":
		getskill(showskill(y))
	case "3":
		getskill(showskill(z))
	default:
		getskill(showskill(roll(10)))
	}
}

func showskill(x int) (string, int) { ////////////////////////////////////////////////////////////////////////顯示能力(待補)
	var s string
	var v int
	switch x {
	case 1:
		s = "攻擊增加"
		v = roll(3)
	case 2:
		s = "防禦增加"
		v = roll(3)
	case 3:
		s = "敏捷增加"
		v = roll(3)
	default:
		s = "血量增加"
		v = roll(5)
	}
	return s, v + 2
}

func getskill(x string, y int) { ///////////////////////////////////////////////////////////////////////////////////////////////獲得能力(待補)
	switch x {
	case "攻擊增加":
		f.Println("攻擊增加", y)
		p0.atk += y
	case "防禦增加":
		f.Println("防禦增加", y)
		p0.def += y
	case "敏捷增加":
		f.Println("敏捷增加", y)
		p0.dex += y
	case "血量增加":
		f.Println("血量增加", y)
		p0.hp += y

	}

}

func mtrlevelup(mtr character) (int, int, int, int, int, int) { /////////怪物升級暫定
	mtr.level += 1
	mtr.hp += 8 + (mtr.hp / 10)
	mtr.atk += 3 + (mtr.atk / 10)
	mtr.def += 3 + (mtr.def / 10)
	mtr.dex += 1 + (mtr.dex / 5)
	mtr.exp += mtr.level
	return mtr.level, mtr.hp, mtr.atk, mtr.def, mtr.dex, mtr.exp
}

func bosslevelup(boss character) (int, int, int, int, int, int) { ////////BOSS升級
	boss.level += 30
	boss.hp += boss.hp
	boss.atk += boss.atk
	boss.def += boss.def
	boss.dex += boss.dex
	boss.exp += boss.exp
	return boss.level, boss.hp, boss.atk, boss.def, boss.dex, boss.exp
}
