---
marp: true
theme: cpslab
paginate: true
header: '-'
footer: '2023/06/05'
---

<!-- _paginate: false -->

# 論文輪講 第 05 回

#### B4 20FI086

#### 橋本慶紀 🦭

---

<!-- _header: '読んだ論文' -->

### Tactile Graphics with a Voice: Using QR Codes to Access Text in Tactile Graphics

学会名: ASSETS '14: Proceedings of the 16th international ACM SIGACCESS conference on Computers &
accessibility • October 2014

---

# 選んだ理由

- 空間音声ラベルプリンタの関連研究として。

---

# 論文概要

- 目の不自由な生徒向けの教科書の図表は、触覚で情報を得られるようになっている。
- しかし、図表に付随するテキストは、目の不自由な生徒にとっていまだアクセシブルではない
- これを解決するために、スマートフォンで QR コードを読み取ることで、触覚グラフィック内のテキストにアクセスするシステムを開発した

---

<!-- _header: '-' -->

## 評価方法

- 被験者に、3 つのタスクを 4 通りの方法で実行してもらい、その結果をアンケート調査する
  - 視覚のみ | 振動モータ | NormalTouch | TextureTouch
  - それぞれで 3 つのタスクをこなしてもらう
  - 正確さやかかった時間も計測

---

<!-- _header: 'この論文を読んで参考になったこと' -->

### 提案そのものについて

- 触覚デバイスの詳細な構造や、使っているセンサの名前がわかって参考になった
- 少人数(今回は 12 人)に幾つかのタスクを試してもらう、という評価手法は参考になる

---

<!-- _header: 'この論文を読んで参考になったこと' -->

### 結果・考察について

- 一見 TextureTouch の方が良さそうに見えるが、実験の結果では、NormalTouch と TextureTouch の間で、精度や満足度に有意差さはないとの事
- 正確な触覚レンダリングができなくても、被験者は正確なように錯覚する
  - 視覚の情報の占める割合がやはり大きい
  - 触覚デバイスはアバウトでもいいのかも
- NormalTouch みたいなの作ろうかと思う
