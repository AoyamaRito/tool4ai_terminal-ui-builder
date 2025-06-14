# ASCII UI Instructor

ASCII UIの作成指示をAIに返すツール

## 概要

このツールは、ASCII形式のUIを作成するための指示を生成します。tool4aiがAIに対して適切な指示を返すことで、ASCII UIとインタラクティブ要素のリストを含む画面を生成できます。

## インストール

```bash
# 1. ファイルをダウンロード
wget https://raw.githubusercontent.com/yourrepo/ascii-ui-instructor.go

# 2. ビルド
go build -o ascii-ui-instructor ascii-ui-instructor.go

# 3. システムにインストール（オプション）
sudo mv ascii-ui-instructor /usr/local/bin/
```

## 使い方

### 基本的な使用（デフォルト：AIへの指示を出力）
```bash
ascii-ui-instructor -layout dashboard
ascii-ui-instructor -layout login
ascii-ui-instructor -layout network
```

### 人間が読める詳細な指示を出力
```bash
ascii-ui-instructor -mode instructions -layout dashboard
```

### 既存のASCII artを分析
```bash
ascii-ui-instructor -mode analyze -input sample.txt
```

## 出力形式

デフォルトモードでは、以下の形式で出力されます：

```
以下のASCII UIを作成してください:

画面名: Login Form
サイズ: 60 x 15 文字

要求仕様:
1. 枠線は +, -, | を使用して描画
2. すべての要素は指定された位置に配置
3. テキストは枠内に収まるように配置

UI要素:
1. Login Title (header)
   位置: center, 1
   内容: LOGIN SYSTEM

[... UI要素のリスト ...]

ASCII UIの下に以下の形式でINTERACTIVE_ELEMENTSセクションを追加:

INTERACTIVE_ELEMENTS:
[Login] - click: Submit login form → login.txt
[Cancel] - click: Cancel login → (画面遷移なし)
username - input: text入力 → (画面遷移なし)
password - input: password入力 → (画面遷移なし)
```

## レイアウトの種類

- `dashboard` - システムモニタリングダッシュボード
- `login` - ログインフォーム
- `network` - ネットワークオペレーションセンター

## カスタマイズ

新しいレイアウトを追加する場合は、`createSampleLayouts()` 関数に新しいレイアウトを追加してください。各UI要素には以下を定義できます：

- Type: 要素の種類（header, box, panel, menu, button等）
- Position: 配置位置
- Size: サイズ
- Content: 表示内容
- Actions: クリックなどのアクション
- Inputs: 入力フィールド

## 画面遷移の記述

INTERACTIVE_ELEMENTSセクションでは、各アクションの結果を記述します：

- `→ filename.txt` - 別の画面に遷移
- `→ (画面遷移なし)` - 同じ画面で処理
- `→ (アプリケーション終了)` - アプリを終了

## ライセンス

MIT License