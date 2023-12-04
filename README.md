# Bike-Festival-2024

## Branch/Commit Type

- feat: 新增/修改功能 (feature)。
- fix: 修補 bug (bug fix)。
- docs: 文件 (documentation)。
- style: 格式 (不影響程式碼運行的變動 white-space, formatting, missing semi colons, etc)。
- refactor: 重構 (既不是新增功能，也不是修補 bug 的程式碼變動)。
- perf: 改善效能 (A code change that improves performance)。
- test: 增加測試 (when adding missing tests)。
- chore: 建構程序或輔助工具的變動 (maintain)。
- revert: 撤銷回覆先前的 commit

### e.g

- feat/auth, refactor/home-layout
- feat: build button component, fix indentation and spacing issues

## dev setup

1. Make sure you have installed `make`, `go` and `docker-compose / docker compose` on your machine
2. Copy `.env.example` to `.env.dev` and modify the environment variables if needed

   ```bash
   cp .env.example .env.dev
   ```

3. Run `make install` to install the tools
   > Run only once when you start the development or the tools required to be updated
4. Run `make dev-up` to start the development database and redis
5. Run `make dev-migrate` to migrate the database
6. Run `make serve` to start the development server with live reload
   1. Webserver will be listening on [localhost:8000](http://localhost:8000), you may change the port in `.env.dev`
7. Run `make dev-down` to stop the development database and redis