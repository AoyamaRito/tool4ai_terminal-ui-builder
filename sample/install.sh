#!/bin/bash

# ASCII UI Instructor インストールスクリプト

echo "ASCII UI Instructor をインストールしています..."

# Goがインストールされているか確認
if ! command -v go &> /dev/null; then
    echo "エラー: Goがインストールされていません。"
    echo "https://golang.org/dl/ からGoをインストールしてください。"
    exit 1
fi

# ビルド
echo "ビルド中..."
go build -o ascii-ui-instructor ascii-ui-instructor.go

if [ $? -ne 0 ]; then
    echo "エラー: ビルドに失敗しました。"
    exit 1
fi

# 実行権限を付与
chmod +x ascii-ui-instructor

# インストール先を選択
echo ""
echo "インストール先を選択してください:"
echo "1) /usr/local/bin (システム全体、要sudo)"
echo "2) ~/bin (現在のユーザーのみ)"
echo "3) 現在のディレクトリ (インストールしない)"
read -p "選択 (1-3): " choice

case $choice in
    1)
        echo "システム全体にインストール中..."
        sudo mv ascii-ui-instructor /usr/local/bin/
        if [ $? -eq 0 ]; then
            echo "インストール完了！"
            echo "使い方: ascii-ui-instructor -layout dashboard"
        else
            echo "エラー: インストールに失敗しました。"
            exit 1
        fi
        ;;
    2)
        echo "ユーザーディレクトリにインストール中..."
        mkdir -p ~/bin
        mv ascii-ui-instructor ~/bin/
        
        # PATHに追加されているか確認
        if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
            echo ""
            echo "注意: ~/bin がPATHに含まれていません。"
            echo "以下を ~/.bashrc または ~/.zshrc に追加してください:"
            echo 'export PATH="$HOME/bin:$PATH"'
        fi
        
        echo "インストール完了！"
        echo "使い方: ascii-ui-instructor -layout dashboard"
        ;;
    3)
        echo "現在のディレクトリに配置しました。"
        echo "使い方: ./ascii-ui-instructor -layout dashboard"
        ;;
    *)
        echo "無効な選択です。"
        exit 1
        ;;
esac

echo ""
echo "使用例:"
echo "  ascii-ui-instructor -layout dashboard    # ダッシュボードのUI指示を生成"
echo "  ascii-ui-instructor -layout login        # ログイン画面のUI指示を生成"
echo "  ascii-ui-instructor -mode instructions -layout dashboard  # 詳細な指示"