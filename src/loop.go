
package main
import(
	"fmt"
	"time"
	"math/rand"
)

func loop_head(){

	choice_cont := []choice{
		choice{"1","初めから"},
		choice{"2","続きから"},
		choice{"3","ストーリー"},
		choice{"4","終了"}};

	for {
		reaction := get_reaction("何をしますか？", choice_cont...)

		curtain()
		
		switch reaction {
		case "1":
			loop_new_game()
			return
		case "2":
			loop_load_game()
		case "3":
			story()
		case "4":
			return
		}
	}
}

func loop_new_game(){

	fmt.Println("新しいゲームを初めます。\n")

	choice_cont := []choice{
		choice{"1","特攻型"},
		choice{"2","バランス型"},
		choice{"3","耐久型"}};
	
	reaction := get_reaction("あなたのGopherのタイプは？", choice_cont...)
	

	var new_gopher_type int
	switch reaction {
	case "1":
		new_gopher_type = -1
		fmt.Println("特攻型ですね。")
	case "2":
		new_gopher_type = 0
		fmt.Println("バランス型ですね。")
	case "3":
		new_gopher_type = 1
		fmt.Println("耐久型ですね。")
	}
	
	curtain()

	fmt.Println("\nそれでは、Gopherの果てしない冒険のはじまりはじまり...\n")
	new_gopher(new_gopher_type)


	curtain()

	empty_input()

	curtain()

	loop_adventure()

}

func loop_load_game(){
	for {

		fmt.Println("\n冒険のcommitIDを入力してください :")
		var input string
		fmt.Scan(&input)

		if load_gopher(input) == true { break }

		fmt.Println("\n無効なIDです")
	}

	curtain()
	get_status()
	loop_adventure()
}

func loop_adventure(){

	

	enemy_rest := 2

	for {

		rand.Seed(time.Now().UnixNano())

		fmt.Printf("ここは %s です。", map_details_cont[adventure_map[gopher.place]])

		if adventure_map[gopher.place] >= 9 {
			village := get_reaction_yes_or_no("入りますか？")
			if village == true {
				loop_village()
			} else {
				curtain()
			}
		} else if gopher.place > 1 && adventure_map[gopher.place] != 8 && adventure_map[gopher.place+1] < 8  && adventure_map[gopher.place-1] < 8 {
			

			if enemy_rest <= 0 {
				loop_fight_random()

				if gopher.gameover == true{
					loop_gameover()
					return
				}
				fmt.Printf("ここは %s です。", map_details_cont[adventure_map[gopher.place]])

				rand.Seed(time.Now().UnixNano())
				enemy_rest = 1 + rand.Intn(5)
			}

		}
		
		choice_cont := []choice{
			choice{"1","北へ進む"},
			choice{"2","南へ進む"},
			choice{"3","ステータスチェック"}};

		reaction := get_reaction("何をしますか？", choice_cont...)

		curtain()
		
		switch reaction {
		case "1":
			gopher.place++
			enemy_rest--
			if gopher.place == len(adventure_map)-1 {
				loop_last_battle()
				return
			}
		case "2":
			if gopher.place != 0 {
				gopher.place--
				enemy_rest--
			} else {
				fmt.Println("後ろには果てしない海が広がっています。進めません...\n...")
			}
		case "3":
			get_status()
		}
	}
}

func loop_village(){

	dweller_num := 0

	curtain()

	fmt.Printf("%s にやってきました！", map_details_cont[adventure_map[gopher.place]])

	for{
		choice_cont := []choice{
			choice{"1","住人と話す"},
			choice{"2","宿屋に行く"},
			choice{"3","教会に行く"},
			choice{"4","ここを出る"}};

		reaction := get_reaction("何をしますか？", choice_cont...)
		curtain()

		switch reaction {
		case "1":
			fmt.Printf("住人%d「%s」\n", dweller_num+1, dwellers_statements[adventure_map[gopher.place]-9][dweller_num])
			dweller_num += 1;
			if dweller_num >= len(dwellers_statements[adventure_map[gopher.place]-9]) { dweller_num = 0 }
		case "2":
			loop_inn()
		case "3":
			loop_church()
		case "4":
			fmt.Printf("%s を抜けました。", map_details_cont[adventure_map[gopher.place]])
			return
		}
	}
}

func loop_inn(){

	owner_is_konqi := adventure_map[gopher.place]%2 == 0

	if owner_is_konqi {
		fmt.Println("Konqi君「やあ！ここは宿屋だよ。」")
	} else {
		fmt.Println("D言語君「どうも。ここは宿屋だよ。」")
	}
	choice_cont := []choice{
		choice{"1","一休みする"},
		choice{"2","ここを出る"}};

	reaction := get_reaction("何をしますか？", choice_cont...)
	curtain()

	switch reaction {
	case "1":
		if owner_is_konqi {
			fmt.Println("Konqi君「OK。じゃあ、ゆっくり休んでね...」\n...\nKonqi「おはよう！今日も頑張って！」")
		} else {
			fmt.Println("D言語君「では、ゆっくりとおやすみください...」\n...\nD言語君「おはよう<3 昨夜はお楽しみでしたね<3」")
		}

		gopher_heal_max()
		fmt.Println("(GopherのHPが全回復した！)")

	case "2":
		if owner_is_konqi {
			fmt.Println("Konqi君「そうか。また来てね。」")
		} else {
			fmt.Println("D言語君「また来なよ。」")
		}
	}
	fmt.Println("\n...")
}

func loop_church(){

	fmt.Println("Octocat司祭「こんにちは。ここは教会です。」\n「あなたの冒険へのcommitを残すことができます。」\n「残しますか？」")
	reaction := get_reaction_yes_or_no("")
	curtain()

	switch reaction {
	case false:
		fmt.Println("Octocat司祭「そうですか。では、さようなら。」\n...")
	case true:
		fmt.Println("Octocat司祭「わかりました。では...」\n")
		fmt.Printf("Octocat司祭「あなたのcommitIDは %s です。メモを忘れないでくださいね。」\n", generate_commit_id(2 + rand.Intn(6)))
		fmt.Println("((ゲーム再開時は、タイトル画面から「続きから」を選択して、このIDを入力してください。))\n...")
	}

}

func loop_fight_random(){
	minlev := -1 
	maxlev := -1
	
	for i, enemy_index := range enemy_cont {
		
		if enemy_index.lev >= (gopher.place/5-1) && minlev == -1 { minlev = i }
		if enemy_index.lev <= (gopher.place+12)/5-1 		     { maxlev = i }
	}

	if minlev >= maxlev || minlev == -1 || maxlev == -1 { return }

	enemy_num := rand.Intn(maxlev-minlev) + minlev

	if(enemy_num >= len(enemy_cont)) { return }

	loop_fight(enemy_num, true)
}

func loop_fight(enemy_num int, escapable bool){
	// gopher.place/5-1, (gopher.place+12)/5-1

	var memo string

	if(gopher.lev >=  enemy_cont[enemy_num].lev+2){
		memo = "倒せそうです。"
	} else if (gopher.lev >= enemy_cont[enemy_num].lev) {
		memo = ""
	} else if (gopher.lev >= enemy_cont[enemy_num].lev-1){
		memo = "強そうです。"
	} else {
		memo = "かなり手強いです！"
	}

	fmt.Printf("\n[[%s が現れた！ %s]]\n", enemy_cont[enemy_num].name, memo)

	enemy_hp := enemy_cont[enemy_num].hp
	looping := true

	for {

		fmt.Printf("\n------------\n%s : %d/%d", enemy_cont[enemy_num].name, enemy_hp, enemy_cont[enemy_num].hp)
		fmt.Printf("\n------------\nGopher (Player) : %d/%d\n\n", gopher_max_hp()-gopher.dmg, gopher_max_hp())

		choice_cont := []choice{
			choice{"1","攻撃する"},
			choice{"2","逃げる"}};
	
		reaction := get_reaction("何をしますか？", choice_cont...)
		curtain()

		

		switch reaction {
		case "1":
			if (gopher.lev > enemy_cont[enemy_num].lev+1-rand.Intn(3)){
				if player_atk(&enemy_num, &enemy_hp) { return }
				if enemy_atk(&enemy_num, &enemy_hp) {
					gopher.gameover = true
					return
				}
			} else {
				if enemy_atk(&enemy_num, &enemy_hp) {
					gopher.gameover = true
					return
				}
				if player_atk(&enemy_num, &enemy_hp) {
					return
				}
			}


		case "2":
			
			if  escapable == false {
				fmt.Println("Gopherは逃げられない！\n...\n")
				break
			}

			fmt.Println("Gopherは必死に走って逃げた！\n...\n")

			if gopher.lev >=  enemy_cont[enemy_num].lev || rand.Intn(enemy_cont[enemy_num].lev-gopher.lev+2) == 0 {
				fmt.Println("うまく逃げ切れたようです。\n")
				looping = false
			} else {
				fmt.Println("追いかけてきました！")
				if enemy_atk(&enemy_num, &enemy_hp) {
					gopher.gameover = true
					return
				}
			}
		}

		if looping == false { break }

	}

}

func player_atk(enemy_num, enemy_hp *int) bool{

	fmt.Println("Gopherの攻撃！\n")

	atk_pw := gopher_atk() + rand.Intn(3)

	fmt.Printf("%sに%dダメージ！\n---\n",  enemy_cont[*enemy_num].name, atk_pw)
	*enemy_hp -= atk_pw

	if *enemy_hp <= 0 {
		fmt.Printf("%sは力尽きました。無事オープンソース化されました。\n", enemy_cont[*enemy_num].name)
		fmt.Printf("Gopherは%d経験値を手に入れました！\n\n", enemy_cont[*enemy_num].exp)
		gopher.exp += enemy_cont[*enemy_num].exp

		level_inc := false

		for {
			if gopher.exp < gopher_exp_need() { break }
			gopher.exp = gopher.exp - gopher_exp_need()
			gopher.lev += 1;
			level_inc = true
		}

		if level_inc == true {
			fmt.Printf("<< Gopherのレベルが上がりました！！ >>\n\n")
			get_status()
		}
		return true
		
	}

	return false
}

func enemy_atk(enemy_num, enemy_hp *int) bool{
	fmt.Printf("%sの攻撃！\n",  enemy_cont[*enemy_num].name)

	atk_pw := enemy_cont[*enemy_num].atk + rand.Intn(3)

	fmt.Printf("Gopherに%dダメージ！\n---\n", atk_pw)
	gopher.dmg += atk_pw
	
	if gopher.dmg >= gopher_max_hp(){
		return true
	}

	return false

}

func loop_gameover(){
	fmt.Printf("ああ...！\nGopherは力尽きました。\nその後、Golangはとうとう有料コンパイラとなってしまいました。\n\n(教会でとったログをもとに、再挑戦してください。)\n")
	curtain()
	fmt.Printf("\nGAME OVER\n\n")
	curtain()
}

func loop_last_battle(){
	fmt.Println("\n???「ふっふっふ....はっはっは。」\n")
	fmt.Println("\n???「お前のような虫けらプロジェクトが、よくここまで来た。相当運がいいようだな。」\n")
	empty_input()
	fmt.Println("\nTux「私の名前はTux。」\n")
	fmt.Println("\nTux「私も元来、お前と同じフリーかつオープンソースな輩であった。」\n")
	fmt.Println("\nTux「しかし、Linus様からいただいたこの強大な商用ライセンスの力を得てから、」\n")
	fmt.Println("\nTux「私はこのような、偉大な姿に生まれ変わったのだ。」\n")
	empty_input()
	fmt.Println("\nTux「お前のような奴に、この強大な力を奪われては困るのだ。」\n")
	fmt.Println("\nTux「さあ。戦おうじゃないか。」\n")
	fmt.Println("\nTux「せめてもの回復ぐらいはしてやろう。フェアな戦いだ。」\n")
	fmt.Println("\nTux「かかってきな！！！スパゲッティコードにしてやるよ！！！」\n")
	empty_input()
	gopher_heal_max()

	// Tux戦
	loop_fight(17, false)

	if gopher.gameover == true {
		loop_gameover()
		return
	}

	empty_input()
	fmt.Println("\nTux「ぐわっ.........何だと.......お前は.......」\n")
	fmt.Println("\nTu?「...........」\n")
	fmt.Println("\nT??「..........！」\n")
	fmt.Println("\n???「.....許さぬ.....」\n")
	fmt.Println("\nなんだ...このプロプライエタリな巨大エネルギーは！！！とGopherが目を疑った、その時だった！！\n")
	empty_input()

	// Linus戦
	loop_fight(18, false)
	if gopher.gameover == true {
		loop_gameover()
		return
	}

	empty_input()
	curtain()
	fmt.Println("\n戦いは終わった。\n")
	fmt.Println("\n空は明るさを取り戻し、太陽の光がGopherの瞳に差した。\n")
	fmt.Println("\nGopherの目に写ったのは、光と氷原の織りなす、見たことのないフリーかつオープンソースな美しい情景であった。\n")
	empty_input()
	fmt.Println("\nGopherの過酷な旅は終わる。フリーかつオープンソースな世界の秩序がまた戻ってくる。\n")
	fmt.Println("\nGopherはFLOSS共和国のフリーかつオープンソースな未来を感じた。\n")
	empty_input()
	fmt.Println("\nさあ！帰ろう！！\n")
	empty_input()
	fmt.Println("\n\n\n\nCreated by Perukii(Tada Teruki) 2020/09/29-2020/09/30\n\nThank you for playing. \n\n")
	curtain()
	fmt.Printf("\nTHE END\n\n")
	curtain()
}