
package main
import(
	"fmt"
	"strconv"
)

// 選択肢クラス
type choice struct {
	key string
	str string
}

func story(){
	fmt.Println(`
	ここは平和なFLOSS共和国。
	Gopherはいつものように街を歩き回っていた。
	しかしその時！
	謎の稲光とともに突然空が闇に包まれ、あたりが急激に混沌としてしまった。
	どうやら、謎の魔法によって、フリーかつオープンソースだったたくさんの善良な国民が
	商用ライセンスにふけてしまい、世の中の平和な秩序が乱れてしまったようだ。

	???「はっはっは。」

	果たしてGopherは、この世界の危機を救えるのか！？
	そして、謎の驚異の正体は、いったい誰なのだろうか...？
	`)
}

func curtain(){
	fmt.Println("=====================================")
}

// 選択肢を提示
func get_reaction(head string, choice_cont ...choice) string{
	
	// 説明文
	fmt.Println(head)

	for{
		// 選択肢の出力
		for _, choice_index := range choice_cont{
			fmt.Printf(" %s)%s\n", choice_index.key, choice_index.str)
		}
		
		fmt.Printf("? ")

		// 入力待ち
		var input string
		fmt.Scan(&input)

		// 入力したキーが選択肢にある？
		// あれば脱却
		for _, choice_index := range choice_cont{
			if input == choice_index.key {

				return choice_index.key
			}
		}

		// なければ再ループ
		fmt.Println("間違った入力です。")

	}

	return ""

}

func get_reaction_yes_or_no(head string) bool{
	choice_cont := []choice{
		choice{"1","はい"},
		choice{"2","いいえ"}};

	reaction := get_reaction(head, choice_cont...)
	result := (reaction == "1")
	
	return result
}

func empty_input(){
	fmt.Println("(続けるには何か文字を入力してください...)")

	// 入力待ち
	var input string
	fmt.Scan(&input)
}

func get_status(){
	fmt.Printf("Gopherのレベルは %d です。\n", gopher.lev)
	fmt.Printf("GopherのHPは %d/%d です。\n", gopher_max_hp()-gopher.dmg, gopher_max_hp())
	fmt.Printf("Gopherの攻撃力は %d です。\n", gopher_atk())
	fmt.Printf("Gopherがレベルアップするにはあと %d 経験値必要です。\n...\n", gopher_exp_need()-gopher.exp)
}

func generate_commit_id(code int) string{

	ityp   := strconv.Itoa(gopher.typ*code)
	ltyp   := strconv.Itoa((len(ityp)+code)%10)
	idmg   := strconv.Itoa(gopher.dmg*code)
	ldmg   := strconv.Itoa((len(idmg)+code)%10)
	iexp   := strconv.Itoa(gopher.exp*code)
	lexp   := strconv.Itoa((len(iexp)+code)%10)
	ilev   := strconv.Itoa(gopher.lev*code)
	llev   := strconv.Itoa((len(ilev)+code)%10)
	iplace := strconv.Itoa(gopher.place*code)
	lplace := strconv.Itoa((len(iplace)+code)%10)
	id := strconv.Itoa(code)+ltyp+ityp+ldmg+idmg+lexp+iexp+llev+ilev+lplace+iplace

	return id
}

