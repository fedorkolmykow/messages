package dbclient

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx"

	m "avito_message/message_service/pkg/modeles"
)

type DbClient interface {
	InsertUser(userAddReq *m.UserAddRequest, createdAt string) (userAddResp *m.UserAddResponse, err error)
	InsertChat(chatAddReq *m.ChatAddRequest, createdAt string) (chatAddResp *m.ChatAddResponse, err error)
	InsertMessage(mesAddReq *m.MessageAddRequest, createdAt string) (mesAddResp *m.MessageAddResponse, err error)
	SelectChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error)
	SelectMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error)
}

type db struct {
	dbCon *pgx.Conn
}

const(
insertUser = "INSERT INTO Users (username, created_at) VALUES ($1, $2) RETURNING user_id;"
insertChat = "INSERT INTO Chats (name, created_at) VALUES ($1, $2) RETURNING chat_id;"
insertChatUsers = "INSERT INTO Chat_Users (chat_id, user_id) VALUES ($1, $2);"
insertMessage = "INSERT INTO Messages (author_id, chat_id, text, created_at " +
	"VALUES ($1, $2, $3, $4) RETURNING message_is);"
selectChats = "SELECT Chats.chat_id, Chats.name, Chats.created_at" +
	" FROM Chats " +
	" LEFT JOIN Chat_Users ON Chat_Users.chat_id = Chats.chat_id" +
	" WHERE Chat_Users.user_id = $1;"
selectMessages = "SELECT Messages.message_id, Messages.author_id, Messages.chat_id, Messages.text, Messages.created_at" +
	"FROM Messages" +
	"WHERE chat_id = $1;"
)

func (d *db) InsertUser(userAddReq *m.UserAddRequest, createdAt string) (userAddResp *m.UserAddResponse, err error) {
	userAddResp = &m.UserAddResponse{}
	row := d.dbCon.QueryRow(context.Background(), insertUser, userAddReq.Username, createdAt)
	err = row.Scan(&userAddResp.UserId)
	if err!=nil{
		fmt.Println(err)
	}
	return
}
func (d *db) InsertChat(chatAddReq *m.ChatAddRequest, createdAt string) (chatAddResp *m.ChatAddResponse, err error) {
	chatAddResp = &m.ChatAddResponse{}
	row := d.dbCon.QueryRow(context.Background(), insertChat, chatAddReq.Name, createdAt)
	err = row.Scan(&chatAddResp.ChatId)
	if err!=nil{
		return
	}

	for i, _ := range chatAddReq.UsersId{
		_, err = d.dbCon.Query(context.Background(), insertChatUsers, id, chatAddReq.UsersId[i])
	}
	return
}
func (d *db) InsertMessage(mesAddReq *m.MessageAddRequest, createdAt string) (mesAddResp *m.MessageAddResponse, err error) {
	mesAddResp = &m.MessageAddResponse{}
	row := d.dbCon.QueryRow(context.Background(), insertMessage,
		mesAddReq.AuthorId, mesAddReq.ChatId, mesAddReq.Text, createdAt)
	err = row.Scan(&mesAddResp.MessageId)
	if err!=nil{
		return
	}

	return
}
func (d *db) SelectChats(chatsGetReq *m.ChatsGetRequest) (chatsGetResp *m.ChatsGetResponse, err error) {
	rows, err := d.dbCon.Query(context.Background(), selectChats, chatsGetReq.UserId)
	defer rows.Close()
	chatsGetResp = &m.ChatsGetResponse{Chats: []m.Chat{}}
	for rows.Next(){
		c := m.Chat{}
		err = rows.Scan(&c.ChatId, &c.Name, &c.CreatedAt)
		if err != nil{
			return
		}
		chatsGetResp.Chats = append(chatsGetResp.Chats, c)
	}

	return
}
func (d *db) SelectMessages(mesGetReq *m.MessagesGetRequest) (mesGetResp *m.MessagesGetResponse, err error) {
	rows, err := d.dbCon.Query(context.Background(), selectMessages, mesGetReq.ChatId)
	defer rows.Close()
	mesGetResp = &m.MessagesGetResponse{Messages: []m.Message{}}
	for rows.Next(){
		m := m.Message{}
		err = rows.Scan(&m.MessageId, &m.AuthorId, &m.ChatId, &m.Text, &m.CreatedAt)
		if err != nil{
			return
		}
		mesGetResp.Messages = append(mesGetResp.Messages, m)
	}
	return
}

// NewDb returns a new Db instance.
func NewDb(ctx context.Context) DbClient {
	connStr := "postgresql://postgres:avitopass@localhost:5555/avitomes"
	dbCon, err := pgx.Connect(ctx, connStr) //os.Getenv("DATABASE_URL")

	if err != nil {
		log.Fatal(err)
	}

	return &db{dbCon: dbCon}
}