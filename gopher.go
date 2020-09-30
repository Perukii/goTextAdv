
package main
import(
	"math"
	"strconv"
)

type player struct{
	typ int // -1:攻撃型 0:バランス型 1:守備型
	exp int
	lev int
	dmg int
	place int
	gameover bool
}

var gopher player

func new_gopher(typ int){
	gopher.typ = typ
	gopher.dmg = 0
	gopher.exp = 0
	gopher.lev = 1
	gopher.place = 0
	gopher.gameover = false
}

//230303032410 

func load_gopher(id string) bool{

	if len(id)<5 { return false }

	i := 1
	j := 0
	var head int
	var val_cont [5]int
	head,_ = strconv.Atoi(string([]rune(id)[0]))

	for {

		var chr,val int
		chr,_ = strconv.Atoi(string([]rune(id)[i]))
		inc := (chr-head+10)%10+1
		val,_ = strconv.Atoi(string([]rune(id)[i+1:i+inc]))

		val_cont[j] = val/head

		i += inc

		if i >= len(id){
			if i == len(id) && j == 4 {
				break
			} else {
				return false
			}
		}
		j++

		if j>4 { return false }
	}

	gopher.typ = val_cont[0]
	gopher.dmg = val_cont[1]
	gopher.exp = val_cont[2]
	gopher.lev = val_cont[3]
	gopher.place = val_cont[4]
	gopher.gameover = false

	return true
}

func gopher_stat(typ int) int{
	switch typ{
	case -1:
		return 60 + gopher.lev*10
	case 0:
		return 52 + gopher.lev*9
	case 1:
		return 48 + gopher.lev*8
	}
	return 0
}

func gopher_atk() int{
	return gopher_stat(gopher.typ)/6 + gopher.lev*2
}

func gopher_exp_need() int{
	return int(math.Pow(1.3, float64(gopher.lev))*10.0)
}

func gopher_max_hp() int{
	return gopher_stat(-gopher.typ)
}

func gopher_heal_max(){
	gopher.dmg = 0
}
