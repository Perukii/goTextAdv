# goTextAdv

Go練習用に作成したテキストアドベンチャーです。<br>
MITライセンス。楽しい！

### ストーリー

ここは平和なFLOSS共和国。<br>
Gopherはいつものように街を歩き回っていた。<br>
しかしその時！<br>
謎の稲光とともに突然空が闇に包まれ、あたりが急激に混沌としてしまった。<br>
どうやら、謎の魔法によって、フリーかつオープンソースだったたくさんの善良な国民が<br>
商用ライセンスにふけてしまい、世の中の平和な秩序が乱れてしまったようだ。<br>
<br>
???「はっはっは。」<br>
<br>
果たしてGopherは、この世界の危機を救えるのか！？<br>
そして、謎の驚異の正体は、いったい誰なのだろうか...？<br>

### あそびかた

プログラムを実行します。
```
go run main.go func.go loop.go gopher.go map.go enemy.go 
```

- はじめての方は「初めから」を選択してゲームを開始します。(1と入力してください)
- プレイヤーは「Gopherの家」から冒険を初めます。
- **プレイヤーは南北の方向にのみ移動可能です。**
- **北に行くほど冒険が進み、敵も強くなります。**
- 「Gopherの家」よりも南に行くことは出来ません。

##### 進めてみる

- とりあえずゲームを開始したら、北に行ってみましょう。
- 途中で敵が出現します。繰り返し攻撃して倒すか、逃げましょう。
- 北に進み続けると、最初の村にたどり着きます。入りましょう。
- 村では**住人と話す**、**宿屋に行く**、**教会に行く**の3つの行動が可能です。
- **住人と話す**を選択すると、村に住む人々のつぶやきを聞くことが出来ます。冒険の上での貴重なヒントを貰うことができるかもしれません。
- **住人と話す**は何度も選択することで、住人が切り替わり、さらに多くのヒントを得ることができます。
- **宿屋に行く**では、HPの回復が出来ます。
- **教会に行く**では、冒険の内容の記録ができ、ゲームを一時中断することができます。
- **HPがなくなると、ゲームオーバーとなり、ゲームが終了します。続けて冒険することはできません。教会でこまめにメモをとっておきましょう。**

##### 便利なTips

- **ステータスチェック**を選択すると、GopherのレベルやHP、経験値などのデータを得ることが出来ます。

##### 続きからはじめる

- ゲームを再開するときは、タイトル画面から「続きから」を選択し、事前に**教会に行く**でメモした内容を改めて入力してください。

##### 最後は...?

- マップの最も北には、世にも恐ろしい存在が待ち受けているようです。
- それは果たして誰なのか、そしてそれを倒した先にあるものは...自分の目(あるいはソースコード)で確かめましょう！

Created by Perukii(Tada Teruki)