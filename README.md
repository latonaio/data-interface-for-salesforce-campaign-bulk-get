# data-interface-for-salesforce-campaign-bulk-get
data-interface-for-salesforce-campaign-bulk-get は、salesforce のキャンペーンオブジェクト取得に必要なデータの整形、および作成時に salesforce から返ってきた response の MySQL への格納をバルク単位で行うマイクロサービスです。

## 動作環境
data-interface-for-salesforce-campaign-bulk-get は、aion-coreのプラットフォーム上での動作を前提としています。  
使用する際は、事前に下記の通りAIONの動作環境を用意してください。  

OS: Linux OS  
CPU: ARM/AMD/Intel  
Kubernetes  
AION のリソース  

## 事前準備  
本マイクロサービスは DB に MySQL を使用します。  
MySQL に関する設定は 環境変数に記述されています。  

| env_name | description |
| --- | --- |
| MYSQL_HOST | ホスト名 |
| MYSQL_PORT | ポート番号 |
| MYSQL_USER | ユーザー名 |
| MYSQL_PASSWORD | パスワード |
| MYSQL_DBNAME | データベース名 |
| MAX_OPEN_CONNECTION | 最大コネクション数 |
| MAX_IDLE_CONNECTION | アイドル状態の最大コネクション数 |
| KANBANADDR: | kanban のアドレス |
| TZ | タイムゾーン |

## kanban との通信
### kanban(ui-backend-for-sharebatakeop) から受信するデータ
kanban から受信する metadata に下記の情報を含む必要があります。

| key | value |
| --- | --- |
| method | 文字列 "get" |
| object | 文字列 "Campaign" |
| connection_type | request |

具体例: 
```example
# metadata (map[string]interface{}) の中身

"method": "get"
"object": "Campaign"
"connection_type": "request"
```

### kanban(salesforce-api-kube) に送信するデータ
kanban に送信する metadata は下記の情報を含みます。

| key | type | description |
| --- | --- | --- |
| method | string | 文字列 "get" を指定 |
| object | string | 文字列 "CampaignReratedList" を指定 |

具体例: 
```example
# metadata (map[string]interface{}) の中身

"method": "get"
"object": "CampaignRelatedList"
```

### kanban(salesforce-api-kube) から受信するデータ
kanban からの受信可能データは下記の形式です


| key | value |
| --- | --- |
| key | 文字列 "CampaignReratedList" |
| content | Campaign の詳細情報を含む JSON 配列|
| connection_type | 文字列 "response" |

具体例:
```example
# metadata (map[string]interface{}) の中身

"key": "CampaignRelatedList"
"content": "[{xxxxxxxxxxxxxx}]"
"connection_type": "response"
```
