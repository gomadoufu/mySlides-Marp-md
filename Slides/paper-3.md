---
marp: true
theme: cpslab
paginate: true
header: '-'
footer: '2023/06/05'
---

<!-- _paginate: false -->

# 論文輪講 & 進捗発表？ 第 03 回

#### B4 20FI086

#### 橋本慶紀 🦭

---

<!-- _paginate: false -->
<!-- _header: '今日話すこと' -->

- **Marp の CSS theme を作ってみた**
- **OC でやったこと**

- **読んだ論文**

---

# OC でやったこと

---

<!-- _header: '読んだ論文 1' -->

### HapticRevolver: Touch, Shear, Texture, and Shape Rendering on a Reconfigurable Virtual Reality Controller

学会名: CHI'18: Proceedings of the 2018 CHI Conference on Human Factors in Computing Systems April 2018

---

# 選んだ理由

- 触覚インタフェース(表示器)を卒研に選ぼうとしている
- Tangible Bits がかなり好きだった。計算機とディスプレイやキーボード以外の方法でインタラクションする方法を提案している論文を探していた
- 最近のものがいい
- 岩井先生が教えてくれた 🌟

---

# 論文概要

- Haptic Revolver という VR コントローラの提案
- 指の下でホイールを回転させることで、触覚フィードバックを得ることができる
- ホイールを交換することで、異なったフィードバックを得ることができる(重要)
- ホイールの速度と方向、素材や形状により、どのような触覚フィードバックを得ることができるかを評価した

---

## この論文の主張

### 提起する問題

- VR 用の HMD とともに使われるコントローラは、触覚の表現が不十分(振動に限定されている)
- 特に、接触刺激が制限されている

---

## この論文の主張

### 先行研究との違い

- 今まで、指に装着するデバイス・手袋型デバイス・ロボットアームなどが提案されていた
- これらは、豊かな触角刺激をもたらすが、装置を身につける必要がある
- 装置を身に付けたくないということで、[ハンディデバイスの研究](https://dl.acm.org/doi/10.1145/2984511.2984526)もある

---

# 紹介動画

[Haptic Revolver](https://dl.acm.org/doi/10.1145/3173574.3173660)

---

<!-- _header: '提案する手法' -->

![w:700](https://raw.githubusercontent.com/gomadoufu/mySlides-Marp-md/image/CPS-Lab/HapticRevolver/HapticRevolver-1.png)

---

<!-- _header: '提案する手法-1' -->

### 仮想環境での位置に応じて昇降や回転を行うホイールを搭載した、ハンディコントローラを作成した

- サーボモータでホイールの昇降
- DC モータ&ロータリーエンコーダでホイール回転
- ホイールは交換可能
- コントローラにはボタンも搭載
- ホイールは 3D プリンタで作成。ホイール上に電子部品を配置した「アクティブホイール」も作成

---

<!-- _header: '提案する手法-1' -->

## ![](https://raw.githubusercontent.com/gomadoufu/mySlides-Marp-md/image/CPS-Lab/HapticRevolver/HapticRevolver-3.png)

---

<!-- _header: '提案する手法-2.1' -->

### VR シーン、手の運動、車輪の種類に応じて、デバイスを適切に制御するレンダリングエンジンを開発

- 指がオブジェクトに近づくにつれ、ホイールを滑らかに上昇させる
- 指が表面を貫通しようとすると、ホイールを強く上昇させ、抵抗を表現
- 指が表面に触れていると、指の動きに合わせてホイールを回転させ、剪断力を表現

---

<!-- _header: '提案する手法-2.2' -->

### ソフトウェア的な工夫

- 指の動きの等倍で車輪を回転させると、すぐにホイール上の面積が足りなくなってしまう。そのため、0.6 倍で回転させる(臨場感を失わない最小の値)
- 指は水平と垂直の差に関して鈍感であるという研究あり
- ホイール記述ファイルを JSON で作成。システムに簡単にホイールを追加できるようにした

---

## ![w:900](https://raw.githubusercontent.com/gomadoufu/mySlides-Marp-md/image/CPS-Lab/HapticRevolver/HapticRevolver-4.png)

---

 <!-- _header: '提案する手法' -->

# 提案手法の優位性

- 装置を身に付けなくてよい。使用者の行動を制限せず、手に取るだけで使い始められる
- さらに、先行研究のハンディデバイスと異なり、指先の剪断力の再現が可能
- そも、感覚の再現ではなく
- 物理的な面に指を接触させれば良いので、トラッキングの必要がない

---

 <!-- _header: '提案する手法' -->

# 評価

全て 5 段階リッカート尺度

**予備実験**

- ホイールの回転方法を変えたときに、違和感があるか

**本実験**

- 振動フィードバックをするコントローラ(HTC VIVE)との比較

---

## ![w:400](https://raw.githubusercontent.com/gomadoufu/mySlides-Marp-md/image/CPS-Lab/HapticRevolver/HapticRevolver-5.png)

---

## ![w:400](https://raw.githubusercontent.com/gomadoufu/mySlides-Marp-md/image/CPS-Lab/HapticRevolver/HapticRevolver-6.png)

---

## ![w:400](https://raw.githubusercontent.com/gomadoufu/mySlides-Marp-md/image/CPS-Lab/HapticRevolver/HapticRevolver-7.png)

---

# 結論

- 実験の結果、Haptic Revolver の方が、振動フィードバックよりも、より自然な触覚フィードバックを提供できることがわかった
- ただし、すべての体験者がリアルなフィードバックを好むわけではなかった

---

# 課題

- ホイールを垂直方向に並べれば、いろんなホイールを交換できるかも
- 手の大きさによって、感触が変わってしまう

---

# まとめ

- **非っ常に面白い論文だった**
- さすがマイクロソフトというか、論文がしっかりしている、、、
- AsiaHaptics 学会面白そう
- 評価、アンケートでいいんだ。リッカート尺度っていうんだ
