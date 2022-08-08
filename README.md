# go_present_slide

社内LT用に作成したスライドです。

## present使用方法

前提としてgoの環境があること。

1. プロジェクトフォルダ作成
1. プロジェクトフォルダ内に拡張子「.slide」ファイル作成
1. フォルダ内で `go mod init` 実行
1. 下記コマンドでインストール
  ```
  go install golang.org/x/tools/cmd/present@latest
  go get golang.org/x/tools/cmd/present
  ```
1. フォルダ内でターミナルに `present` を入力しpresentを開始
1. 下記のような表示が出るので、「127.0.0.1:XXXXX」をクリック
  `2022/MM/DD HH:MM:SS accepting connection from: 127.0.0.1:XXXXX`
1. ブラウザが開かれるので、作成した「.slide」をクリックするとスライドが表示される
