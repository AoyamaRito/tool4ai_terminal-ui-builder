# Tool4AI Terminal UI Builder

🤖 **AI優先ターミナルインターフェースデザイナー**

📖 **[日本語版 README](README_JA.md) | [English README](README.md)**

**📋 Claude Codeが完璧にアライメントされたアスキーアートUI設計図を瞬時に作成可能！**

**Claude Code、ChatGPT、その他のAI開発ツール専用に設計されています。**

**完璧なワークフロー: AI作成 → 人間確認 → AI実装**

1. 🤖 **AIが簡単にアスキーアートUIを作成** - Claude Codeが完璧なアライメントで複雑なインターフェースを生成
2. 👁️ **人間が特別なビューワー無しで閲覧可能** - どこでも読めるプレーンテキスト形式
3. 💻 **AIがこの設計図からコーディング可能** - Claude Codeがアスキー設計図を実装の設計図として使用

このツールを使用することで、Claude Codeはアライメント問題なしにプロフェッショナルで正確なアスキーアート ターミナルインターフェースを生成できます。手動でのアスキーアート間隔問題に苦労する代わりに、Claude Codeは単一のコマンドで完璧な精度で複雑なダッシュボード、フォーム、UIを作成できます。

## 🎯 完全な設計からコーディングまでのワークフロー

**AIアシスタント向け:**
- **🤖 簡単なアスキーUI作成** - アライメントに苦労することなく複雑なインターフェースを生成
- **📐 完璧なアスキーアートアライメント** - 毎回正確にアライメントされたターミナルUI
- **💻 コーディングの設計図** - 生成された設計を実装ガイドとして使用
- **⚡ 瞬時の生成** - シンプルなコマンドでプロフェッショナルなインターフェースを作成

**人間の開発者向け:**
- **👁️ ユニバーサル表示** - 特別なビューワー不要、任意のテキストエディタで読める
- **📋 明確な設計文書** - 仕様書やドキュメンテーションに最適
- **🤝 AI-人間コラボレーション** - AIが設計作成、人間がレビュー、AIが実装
- **✅ 一貫したコミュニケーション** - 設計・開発フェーズで同じフォーマット

![ASCII Art](https://img.shields.io/badge/ASCII-Art-brightgreen)
![Go](https://img.shields.io/badge/Go-1.19+-blue)
![Cross Platform](https://img.shields.io/badge/Platform-Cross--Platform-orange)
![License](https://img.shields.io/badge/License-MIT-green)

## ✨ 機能

- 🚀 **5つのハイパーパーフェクト プロフェッショナルテンプレート**
  - ネットワークオペレーションセンター
  - トレーディングフロア ターミナル
  - サイバーディフェンス コマンドセンター
  - データコントロール ステーション
  - スペースミッション コントロール

- 🔧 **コア機能**
  - リアルタイム品質チェック
  - 座標付きグリッドビュー
  - 安全なテキスト置換（構造を破壊しない）
  - クロスプラットフォーム シングルバイナリ
  - 一貫性のための80文字固定幅

- 🛡️ **安全機能**
  - スペースのみ置換 - 構造を上書きしない
  - テキストが長すぎる場合の自動切り詰め
  - 一般的な単語のインテリジェント略語
  - ASCII専用文字検証

## 🚀 クイックスタート

### AIアシスタント向け

**Claude Code、ChatGPT、その他のAIツールに以下のようなプロンプトで指示するだけ:**

- *「terminal-ui-builderを使ってネットワークオペレーションダッシュボードを作成してください」*
- *「tradingレイアウトを使ってトレーディングインターフェースを生成してください」*
- *「このツールでアスキーアートのログインフォームを作成してください」*
- *「terminal-ui-builderを使って私のアスキーアートファイルのアライメントをチェックしてください」*

**AIツールは直接これらのコマンドを実行できます:**

```bash
# AIがプロフェッショナルなダッシュボードを瞬時に生成
./terminal-ui-builder -layout netops -h 25

# 金融アプリ用トレーディングインターフェースを作成
./terminal-ui-builder -layout trading -h 30

# ログインフォームを生成
./terminal-ui-builder -layout login -h 18

# AIが生成したアスキーアートを検証
./terminal-ui-builder -mode check -f user_interface.txt
```

### 人間の開発者向け

#### macOS
```bash
# Intel Mac
curl -L https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-darwin-amd64 -o terminal-ui-builder
chmod +x terminal-ui-builder

# Apple Silicon Mac
curl -L https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-darwin-arm64 -o terminal-ui-builder
chmod +x terminal-ui-builder
```

#### Linux
```bash
# AMD64
curl -L https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-linux-amd64 -o terminal-ui-builder
chmod +x terminal-ui-builder

# ARM64
curl -L https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-linux-arm64 -o terminal-ui-builder
chmod +x terminal-ui-builder
```

#### Windows
```powershell
# PowerShell - AMD64
Invoke-WebRequest https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-windows-amd64.exe -OutFile terminal-ui-builder.exe

# PowerShell - ARM64
Invoke-WebRequest https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-windows-arm64.exe -OutFile terminal-ui-builder.exe
```

#### ソースからビルド
```bash
git clone https://github.com/AoyamaRito/tool4ai_terminal-ui-builder.git
cd tool4ai_terminal-ui-builder
go build -o terminal-ui-builder terminal-ui-builder.go
```

### AI向けコマンド

AIアシスタントはこれらのシンプルなパターンを使用できます:

```bash
# 特定のレイアウトを生成
./terminal-ui-builder -layout [netops|trading|cyber|data|space] -h [height]

# 任意のアスキーアートの品質チェック
./terminal-ui-builder -mode check -f [filename]

# グリッドオーバーレイでデバッグ
./terminal-ui-builder -mode view -f [filename]
```

## 🎯 利用可能レイアウト

### プロフェッショナルテンプレート

| レイアウト | 説明 | 用途 |
|--------|-------------|----------|
| `netops` | ネットワークオペレーションセンター | インフラ監視 |
| `trading` | トレーディングフロア ターミナル | 金融アプリケーション |
| `cyber` | サイバーディフェンス コマンド | セキュリティオペレーション |
| `data` | データコントロール ステーション | ビッグデータ管理 |
| `space` | スペースミッション コントロール | 科学アプリケーション |

### 基本テンプレート

| レイアウト | 説明 | 用途 |
|--------|-------------|----------|
| `dashboard` | シンプルな管理ダッシュボード | 汎用 |
| `login` | ログインフォーム | 認証 |
| `box` | タイトル付きカスタムボックス | 汎用コンテナ |

## 📋 コマンドリファレンス

### モード

- **create** - 新しいアスキーアート インターフェースを生成
- **check** - 既存のアスキーアートのアライメント問題を検証
- **view** - デバッグ用の座標グリッドで表示

### オプション

```
-layout string    レイアウトタイプ (デフォルト "netops")
-h int           キャンバスの高さ (デフォルト 25)
-w int           ボックスレイアウトの幅 (デフォルト 40, 最大 80)
-t string        レイアウトのタイトル
-f string        チェックまたは表示するファイル
-mode string     モード: create, check, view (デフォルト "create")
```

## 🌟 ライブデモ

### 🚀 複雑なネットワークオペレーションセンター

AIが瞬時に高度な監視ダッシュボードを生成できます:

```bash
./terminal-ui-builder -layout netops -h 35
```

```
+------------------------------------------------------------------------------+
| NetOps CmdCtr                                        [LIVE] [ALERT] [X]      |
+------------------------------------------------------------------------------+
|                                                                              |
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| | STATUS   | | CPU      | | MEMORY   | | NETWORK  | | STORAGE  | | LATENCY  ||
| | CORE     | | 73%      | | 12GB/32GB| | 89%      | | 2.1TB/5TB| | 12ms     ||
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
|                                                                              |
| +----------------------------+  +------------------------------------------+ |
| | LIVE TRAFFIC FLOW          |  | INCIDENT RESPONSE                        | |
| | US-WEST [====] 892 Mbps    |  | HIGH: SSL Cert Expiring 3d               | |
| | US-EAST [===] 634 Mbps     |  | MED: High mem usage DB-02                | |
| | EU-WEST [==] 423 Mbps      |  | LOW: CDN cache miss >15%                 | |
| +----------------------------+  +------------------------------------------+ |
| Ops: 4 engineers | Global: 99.97% uptime                                     |
+------------------------------------------------------------------------------+
```

### 💹 プロフェッショナル トレーディング ターミナル

金融AI アプリケーションに最適:

```bash
./terminal-ui-builder -layout trading -h 30
```

### 🛡️ サイバーセキュリティ コマンドセンター

AI生成セキュリティオペレーション ダッシュボード:

```bash
./terminal-ui-builder -layout cyber -h 25
```

### 🚀 スペースミッション コントロール

科学・研究AI アプリケーション:

```bash
./terminal-ui-builder -layout space -h 25
```

### 🧪 API テスティング スタジオ

プロフェッショナルなAPI開発・テスト インターフェース:

（実際のAPIテスティングスタジオのアスキーアートをここに表示）

### ⚡ ワンコマンド生成

**AIアシスタントは単一のコマンドでこれらの複雑なインターフェースを作成できます:**

```bash
# ネットワークオペレーション ダッシュボード
./terminal-ui-builder -layout netops -h 35 > network_dashboard.txt

# トレーディングフロア ターミナル
./terminal-ui-builder -layout trading -h 30 > trading_terminal.txt

# セキュリティオペレーション センター
./terminal-ui-builder -layout cyber -h 25 > security_center.txt
```

**最適な用途:**
- 🤖 AIドキュメント生成
- 📋 システム設計議論
- 🎓 インターフェース概念の教育
- 🔧 高速プロトタイピング セッション

## 🔍 品質チェック

ツールには包括的な品質チェックが含まれています:

```bash
./terminal-ui-builder -mode check -f myart.txt
```

- ✅ 垂直線アライメント検証
- ⚠️ 非ASCII文字検出
- 📍 正確なエラー位置報告
- 🛡️ 構造整合性検証

## 📁 プロジェクト構造

```
tool4ai_terminal-ui-builder/
├── terminal-ui-builder.go  # メインプログラム
├── examples/              # サンプル出力
├── docs/                  # ドキュメント
└── releases/              # プリビルドバイナリ
```

## 🤝 貢献

1. リポジトリをフォーク
2. フィーチャーブランチを作成
3. 変更を行う
4. 新しいレイアウトのテストを追加
5. プルリクエストを送信

## 📜 ライセンス

MIT License - 詳細はLICENSEファイルを参照してください。

## 🎯 ユースケース

**AIアシスタントに最適:**
- 🤖 **瞬時のUI生成** - AIが単一のコマンドで複雑なインターフェースを作成
- 📝 **ドキュメント支援** - AIが仕様書用の視覚的モックアップを生成
- 🎓 **教育ツール** - AIが実際の例でUI概念を実演
- 🔧 **高速プロトタイピング** - AIが開発議論中にUIモックアップを構築
- 📊 **システム設計** - AIがアーキテクチャ図やダッシュボードを作成

**人間の開発者に最適:**
- 📊 システム監視ダッシュボード
- 💹 金融アプリケーション インターフェース
- 🛡️ セキュリティオペレーション センター
- 🚀 科学ミッション コントロール
- 📚 技術ドキュメント

## 🔗 関連プロジェクト

Tool4AI エコシステムの一部 - AI開発のためのより良いツールの構築。

---

## 🤖 AI統合例

**Claude Code 使用法:**
```bash
# AIがネットワークオペレーション ダッシュボードを瞬時に作成
./terminal-ui-builder -layout netops -h 25 > network_dashboard.txt

# AIがアスキーアート アライメント問題を検証・修正
./terminal-ui-builder -mode check -f my_interface.txt
```

**ChatGPT 統合:**
```bash
# プロフェッショナルなトレーディング インターフェースを生成
./terminal-ui-builder -layout trading -h 30

# スペースミッション コントロール インターフェースを作成
./terminal-ui-builder -layout space -h 25
```

**AIワークフローに最適:**
- 手動アスキーアートの苦労なし
- 一貫したプロフェッショナルな出力
- 瞬時の検証とデバッグ
- クロスプラットフォーム互換性

---

**AIアシスタントとAI開発コミュニティのために ❤️ で構築**