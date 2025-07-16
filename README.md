# GoのgRPCによるDDD構成のAPIサンプル
Go言語（Golang）のgRPCおよびDDD（ドメイン駆動設計）によるバックエンドAPI開発用サンプルです。  
  
<br />
  
## DDDのディレクトリ構成　　
ディレクトリ構成としてはDDDの思想に基づいたレイヤードアーキテクチャを採用しています。  
  
```
/src
├── /internal
|   ├── /application（アプリケーション層）
|   |    └── usecase（ユースケース層）
|   |
|   ├── /domain（ドメイン層）
|   |    ├── model（ドメインモデルの定義。ビジネスロジックは可能な限りドメインに集約させる。）
|   |    ├── （仮）repository（リポジトリのインターフェース定義）
|   |    └── （仮）service（外部サービスのインターフェース定義）
|   |
|   ├── /infrastructure（インフラストラクチャ層）
|   |    ├── （仮）database（データベース設定）
|   |    ├── logger（ロガーの実装。インターフェース部分はユースケース層で定義。）
|   |    ├── （仮）persistence（リポジトリの実装。DB操作による永続化層。）
|   |    ├── （仮）cache（キャッシュを含めたリポジトリの実装。インターフェースはリポジトリと同一。）
|   |    └── （仮）externalapi（外部サービスの実装）
|   |
|   ├── /presentation（プレゼンテーション層）
|   |    ├── server（サーバー層）
|   |    ├── interceptor（インターセプターの定義）
|   |    └── router（ルーター設定。レジストリのコントローラーを利用して設定する。）
|   |
|   └── /registry（レジストリ層。依存注入によるサーバーのインスタンスをコントローラーにまとめる。）
|
├── /pb（protoファイルから生成したProtocol Buffersのコード）
|
└── /proto（スキーマ定義）
```
> <span style="color:red">※（仮）のものは将来的に追加する想定の例</span>  
  
</br>
  
### APIの作成手順  
  1. ドメインの定義  
    ドメインを新規追加、または既存のドメインにビジネスロジックの追加。  
    永続化が必要ならリポジトリの定義、外部サービスとの連携が必要ならサービスの定義を追加。 
  
  2. リポジトリやサービスの実装  
    リポジトリやサービスのインターフェース定義を追加した場合、インフラストラクチャ層に実装を定義。  
  
  3. スキーマ定義  
    protoファイルでAPIのスキーマを定義。  
  
  4. Protocol Buffersのコード生成  
    Bufを利用し、protoファイルからProtocol Buffersのコードを生成。  
  
  5. ユースケースの定義  
    Protocol Buffersのスキーマ定義およびドメインやリポジトリなどを用いて、ユースケースにビジネスロジックを定義。  
  
  6. サーバーの定義  
    ユースケースを用いてサーバーの定義。  
  
  7. レジストリ登録  
    リポジトリ、ユースケース、サーバーのインスタンスをレジストリのコントローラーに登録。  
  
  8. ルーター設定の追加  
    レジストリを用いてルーター設定を追加。
  
<br />
  
## 要件
・Goのバージョンは<span style="color:green">1.24.x</span>です。  
  
<br />
  
## ローカル開発環境構築
### 1. 環境変数ファイルをリネーム
```
cp ./src/.env.example ./src/.env
```  
  
### 2. コンテナのビルドと起動
```
docker compose build --no-cache
docker compose up -d
```  
  
### 3. コンテナの停止・削除
```
docker compose down
```  
  
<br />
  
## コード修正後に使うコマンド
ローカルサーバー起動中に以下のコマンドを実行可能です。  
  
### 1. go.modの修正
```
docker compose exec grpc go mod tidy
```  
  
### 2. フォーマット修正
```
docker compose exec grpc go fmt ./internal/...
```  
  
### 3. コード解析チェック
```
docker compose exec grpc staticcheck ./internal/...
```  
  
### 4. モック用ファイル作成（例）  
・リポジトリのモックファイル作成
```
docker compose exec grpc mockgen -source=./internal/domain/XXX/XXX_repository.go -destination=./internal/domain/XXX/mock_XXX_repository/mock_XXX_repository.go
```  
  
・ユースケースのモックファイル作成  
```
docker compose exec grpc mockgen -source=./internal/application/usecase/XXX/XXX.go -destination=./internal/application/usecase/XXX/mock_XXX/mock_XXX.go
```
  
### 5. テストコードの実行
・テストコードのファイル（ _test.go ）を追加したパッケージのみテストを実行
```
docker compose exec grpc go test -v $(docker compose exec grpc go list -f '{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}' ./...)
```  
> ※オプション「-cover」を付けるとカバレッジも確認できます。カバレッジは80%以上推薦です。  
  
### 6. テストコードのカバレッジ対象確認用のファイル出力
必要に応じて以下のコマンドを実行し、出力されるファイルからカバレッジ対象のコードを確認して下さい。  
```
docker compose exec grpc go test -v -coverprofile=internal/coverage.out $(docker compose exec grpc go list -f '{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}' ./...)

docker compose exec grpc go tool cover -html=internal/coverage.out -o=internal/coverage.html
```  
> <span style="color:red">※src/internal/coverage.htmlをブラウザで開いて確認して下さい。</span>  
  
<br />
  
## .protoファイルからProtocol Buffersのコード生成
ローカルサーバー起動中に以下のコマンドを実行可能です。  
  
・buf.yamlから依存関係のインストール（buf.lockを生成）  
```
docker compose exec grpc buf dep update
```  
  
・buf.gen.yamlからpbファイルを生成
```
docker compose exec grpc buf generate
```  
  
<br />
  
## 参考記事  
[]()  
  