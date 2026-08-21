```markdown
# Go Blog Backend API

這是一個使用 **Go**、**Gin** 和 **GORM** 開發的 RESTful 部落格後端系統，具備使用者認證（JWT）、密碼雜湊加密（Bcrypt），以及完整的文章與評論 CRUD（增刪改查）功能。

## 專案功能

- **使用者認證**：使用 JWT (JSON Web Token) 與 Bcrypt 實現安全的使用者註冊與登入。
- **文章管理**：已認證的使用者可以建立、更新和刪除自己的部落格文章；任何人都可以檢視所有文章或單篇文章詳情。
- **評論系統**：已認證的使用者可以對文章發表評論，且所有人都能檢視特定文章的所有評論。
- **資料庫 ORM**：結合 GORM 與 SQLite，實現自動資料模型遷移與關聯式資料管理。

---

## 執行環境與需求

請確保你的電腦上已安裝以下工具：
- **Go**（建議版本 1.18 或以上）
- **Git**（選填）

---

## 安裝與相依套件設定

1. 將專案下載至本地，並進入專案根目錄：
   ```bash
   cd blog-backend

```

2. 執行以下指令自動下載並整理所有需要的相依套件（Gin、GORM、SQLite、JWT、Bcrypt）：
```bash
go mod tidy

```



---

## 如何啟動專案

1. 啟動後端伺服器：
```bash
go run main.go

```


2. 伺服器將會在 **`http://localhost:8080`** 啟動並開始監聽。
*(註：啟動時會在根目錄自動產生一個名為 `blog.db` 的 SQLite 資料庫檔案)*。

---

## API 介面參考指南

### 1. 公開介面（無需驗證）

* **註冊使用者**：`POST /register`
* Body 參數：`{"username": "alice", "password": "password123", "email": "alice@example.com"}`


* **登入取得 Token**：`POST /login`
* Body 參數：`{"username": "alice", "password": "password123"}`
* *成功後會回傳一串 JWT `token`。*


* **取得所有文章**：`GET /posts`
* **取得單篇文章**：`GET /posts/:id`
* **取得文章評論**：`GET /posts/:id/comments`

### 2. 受保護介面（需要 JWT 驗證）

*註：訪問這些介面時，請在 HTTP Request Header 中帶入 `Authorization: Bearer <你的_token>*`

* **建立文章**：`POST /posts`
* Body 參數：`{"title": "My First Post", "content": "Hello World!"}`


* **更新文章**：`PUT /posts/:id` *（僅限文章作者）*
* Body 參數：`{"title": "Updated Title", "content": "Updated content."}`


* **刪除文章**：`DELETE /posts/:id` *（僅限文章作者）*
* **發表評論**：`POST /comments`
* Body 參數：`{"content": "Great article!", "post_id": 1}`



```

```
