FROM golang:1.20-alpine
# 有設定WORKDIR就不需要mkdir，會自己創資料夾
WORKDIR /app
# COPY ./src ./dest
COPY . .
# RUN 用多個RUN去寫的話，會產生新的image，因此最好是使用一次RUN，然後換行處理
# go build的時候中過一招卡很久，沒有指定-o輸出檔名的話，會使用資料夾預設的名稱當做app name，然後ENTRYPOINT怎樣都找不到
RUN go mod tidy && \
    go build -o myGolangApp
# ENTRYPOINT docker執行的進入點，預設工作目錄是/app，所以才可以用相對路徑直接找到myGolangApp進入
ENTRYPOINT ./myGolangApp