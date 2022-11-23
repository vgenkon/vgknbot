package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"upgrade/db"
	_ "upgrade/db"
)

type Client struct {
	bot *tgBotApi.BotAPI
}

func New(apiKey string) *Client {
	bot, err := tgBotApi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}
	return &Client{
		bot: bot,
	}
}

func (c *Client) SendMessage(text string, chatId int64) error {
	msg := tgBotApi.NewMessage(chatId, text)
	msg.ParseMode = "Markdown"
	_, err := c.bot.Send(msg)
	return err
}

type usr struct {
	ID         int
	Name       string
	TelegramID int64
	FirstName  string
	LastName   string
}

func Sender(id int64) {
	db.GetMsg()
	msg := *db.MsgOutPointer
	c := New("5964396245:AAGqN6irdMcDRzKeWDh9GsLp8Rh8y3JtFIM")
	err := c.SendMessage(msg, id)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("Сообщение отпралено пользователю: %v\n", id)
}

func main() {
	dbs, err := sql.Open("sqlite3", "upgrade.db")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer func(dbs *sql.DB) {
		err := dbs.Close()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}(dbs)
	rows, err := dbs.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}(rows)
	var users []usr
	for rows.Next() {
		u := usr{}
		err := rows.Scan(&u.ID, &u.Name, &u.TelegramID, &u.FirstName, &u.LastName)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		users = append(users, u)
	}
	Sl := make([]int64, 0)
	for _, val := range users {
		Sl = append(Sl, val.TelegramID)
	}
	for i := 0; i < len(Sl); i++ {
		Sender(Sl[i])
	}
}
