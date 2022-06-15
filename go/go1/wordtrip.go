package main

import (
	f "fmt"
)

func main() {
	title()
}

func title() { //////////////開始畫面
	f.Println("歡迎來到文字之旅~")
	f.Println("輸入數字開始你的旅程!")
	f.Println("1.開始新遊戲 2.繼續遊戲 3.結束遊戲")
	switch userscan() {
	case "1":
		f.Println("你感到一陣天旋地轉~")
		resetall()
		gamestart()
	case "2":
		f.Println("你睜開了眼睛。") /////////////////////////////////////////////////////////////////////////////////////////////////紀錄(還沒有)
		InMap()
	case "3":
		f.Println("後會有期了。") /////////////////結束
	default:
		f.Println("沒有這個選項喔!")
		title()
	}
}

func gamestart() { //遊戲開始
	f.Println("你睜開了眼睛，來發現自己來到了一個陌生的世界。")
	f.Println("你面前出現了一位女神。")
	f.Println("女神:這個世界正在被魔族侵略，")
	f.Println("請拯救他們吧!")
	var x int8 = 0
	for x <= 10 && x >= 0 {
		f.Println("1.沒問題，交給我吧 2.這跟我無關，我不要。 3.轉身離開。")
		switch userscan() {
		case "1":
			f.Println("出發吧勇士!!")
			x -= 10
		case "2":
			f.Println("請拯救他們吧!")
			x++
		case "3":
			f.Println("眼前一黑，你又回到了女神面前。")
			x++
		default:
			f.Println("請回應女神的請求。") /////////////////////////無法拒絕
			x++
		}
	}
	if x < 10 {
		letsgo()
	} else { ////////////////拒絕10次死亡
		dead()
	}
}

func letsgo() { //出發
	f.Println("女神:在出發之前，這裡有三件裝備，")
	f.Println("你可以在裡面選擇其中一件，")
	f.Println("但是裝備的好壞只能憑你的運氣了。")
	var a, x, y, z int ////////////////////次數 攻擊 防禦 敏捷

	for a < 3 {
		f.Println("1.大地之劍 2.森林之鎧 3.風之長靴")
		switch userscan() {
		case "1":
			for x < 5 || x > 10 {
				x = roll(10)
				a = 6
			}
			item.atk = x
			getitem(itemlevel(x), "大地之劍")
		case "2":
			for y < 5 || y > 10 {
				y = roll(10)
				a = 6
			}
			item.def = y
			getitem(itemlevel(y), "森林之鎧")
		case "3":
			for z < 5 || z > 10 {
				z = roll(10)
				a = 6
			}
			item.dex = z
			getitem(itemlevel(z), "風之長靴")
		default:
			f.Println("只有這三件可以選喔~")
			a += 1
			continue
		}
	}
	if a == 3 {
		f.Println("真是貪心的勇者，那就都拿去吧!") //////////////三項數值固定
		x = 6
		item.atk = x
		y = 6
		item.def = y
		z = 6
		item.dex = z
		getitem(itemlevel(x), "大地之劍")
		getitem(itemlevel(y), "森林之鎧")
		getitem(itemlevel(z), "風之長靴")

	}
	f.Println("帶上裝備開始你的旅程吧!")
	InMap()
}

func InMap() { //////////////////////////城內
	f.Println("你眼睛一眨，發現自己站在街道上!")
	f.Println("1.出城冒險 2.魔法商店 3.冒險公會")
	switch userscan() {
	case "1":
		outside()
	case "2":
		shop()
	case "3":
		groap()
	case "4":
		break
	default:
		InMap()
	}
}

func outside() { ////////////////出城
	f.Println("你來到了城外，現在你要做什麼呢?")
	f.Println("1.戰鬥 2.休息一下 3.回城")
	switch userscan() {
	case "1":
		fight()
		outside()
	case "2":
		if p1.hp < p0.hp {
			p1.hp += p0.hp / 2 /////////////////////////////////////////////////暫
		}
		outside()
	case "3":
		InMap()
	default:
		outside()
	}
}

func shop() { ////////////////////////////商店
	f.Println("嘿嘿嘿~歡迎來到魔法商店~")
	f.Println("你可以用10點體力來進行道具抽選~")
	f.Printf("你現在有%d點體力，要試試手氣嗎~\n", p1.hp)
	f.Println("1.來吧!!讓你看看歐洲之手!! 2.我只是進來看看而已。(離開)")
	switch userscan() {
	case "1":
		if p1.hp > 10 {
			f.Println("你面前憑空打開了一個盒子!")
			chooseitem(p1.level)
			p1.hp -= 10
		} else {
			f.Println("哎呀~看來你的體力不夠呢~")
			InMap()
		}
	case "2":
		InMap()
	default:
		f.Println("看來你對這樣的選擇並不滿足呢~")
		f.Println("但你的資格夠嗎~嘿嘿嘿~")
		if p1.hp > 30 && p1.level > 15 {
			f.Println("看來你有這個資格呢~")
			f.Println("那就看看這些吧~")
			chooseitem(p1.level * 2)
		} else {
			f.Println("資格不夠呢~~嘿嘿~")
			InMap()
		}
	}
}

func groap() { /////////////////////////公會
	x, y, z := roll3(5)
	f.Println("你來到冒險者公會的任務欄前。")
	f.Println("選擇你要狩獵的目標:")
	f.Printf("1.%s 2.%s 3.%s 4.離開公會\n", questchoose(x), questchoose(y), questchoose(z))
	switch userscan() {
	case "1":
		qusetfight(questchoose(x))
	case "2":
		qusetfight(questchoose(y))
	case "3":
		qusetfight(questchoose(z))
	case "4":
		InMap()
	default:
		qusetfight(questchoose(roll(5)))
	}
}

func endgame() { //////////////////完結
	updatep1()
	f.Println("你打敗了魔龍，為這個世界帶來了和平。")
	f.Println("要繼續冒險嗎?")
	f.Println("1.繼續冒險 2.結束冒險")
	f.Println(p1)
	switch userscan() {
	case "1":
		f.Println("魔龍不知何時會在捲土重來，你必須再強化自己!")
		f.Println("為下一次的戰鬥做好準備。")
		InMap()
	case "2":
		break
	default:
		f.Println("黑暗中，張開了一雙血紅的眼睛......(待續)") //////////////////////////////////////////////////////////////////////////////////////////////////////
		title()
	}

}

func dead() { /////////////////死亡
	f.Println("YOU DEAD")
	title()
}
