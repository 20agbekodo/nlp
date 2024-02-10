package main

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("mysql", "admin:admin@/gin_react")
    if err != nil {
        panic(err)
    }
}

type User struct {
    ID uint `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Conversation struct {
    ID uint `json:"id"`
    UserID uint   `json:"user_id"`
    Title  string `json:"title"`
}

type Message struct {
    ID uint `json:"id"`
    Content  string `json:"content"`
    Date  time.Time `json:"date"`
    ConversationID uint   `json:"conversation_id"`
    IsUser  bool `json:"is_user"`
}

func Register(c *gin.Context) {
    var newUser User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", newUser.Username, newUser.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    id, err := result.LastInsertId()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    newUser.ID = uint(id)

    c.JSON(http.StatusOK, newUser)
}

func Login(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    row := db.QueryRow("SELECT * FROM user WHERE username = ? AND password = ?", user.Username, user.Password)
    var foundUser User
    err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, foundUser)
}

func GetUsers(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM user")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Username, &user.Password) // Update the Scan function call
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        users = append(users, user)
    }

    c.JSON(http.StatusOK, users)
}

func CreateConversation(c *gin.Context) {
    var newConversation Conversation
    if err := c.ShouldBindJSON(&newConversation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    res, err := db.Exec("INSERT INTO conversation (user_id, title) VALUES (?, ?)", newConversation.UserID, newConversation.Title)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    id, err := res.LastInsertId()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    newConversation.ID = uint(int(id))
    c.JSON(http.StatusOK, newConversation)
}

func GetConversations(c *gin.Context) {
    userID := c.Query("user_id")
    rows, err := db.Query("SELECT * FROM conversation WHERE user_id = ?", userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var conversations []Conversation
    for rows.Next() {
        var conversation Conversation
        err := rows.Scan(&conversation.ID, &conversation.UserID, &conversation.Title)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        conversations = append(conversations, conversation)
    }

    c.JSON(http.StatusOK, conversations)
    defer rows.Close()
}

func DeleteConversation(c *gin.Context) {
    var deleteRequest struct {
        ConversationID uint `json:"conversation_id"`
    }
    if err := c.ShouldBindJSON(&deleteRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.Exec("DELETE FROM conversation WHERE id = ?", deleteRequest.ConversationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Conversation deleted successfully"})
}

func CreateMessage(c *gin.Context) {
    var newMessage Message
    if err := c.ShouldBindJSON(&newMessage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    res, err := db.Exec("INSERT INTO message (conversation_id, content, is_user) VALUES (?, ?, ?)", newMessage.ConversationID, newMessage.Content, newMessage.IsUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    id, err := res.LastInsertId()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    newMessage.ID = uint(int(id))
    c.JSON(http.StatusOK, newMessage)
}

func GetConversationMessages(c *gin.Context) {
    conversationID := c.Query("conversation_id")

    rows, err := db.Query("SELECT * FROM message WHERE conversation_id = ?", conversationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var messages []Message
    for rows.Next() {
        var message Message


        var date string
        
        err := rows.Scan(&message.ID, &message.Content, &date, &message.ConversationID, &message.IsUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        message.Date, err = time.Parse("2006-01-02 15:04:05", date)
                if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        messages = append(messages, message)
    }

    c.JSON(http.StatusOK, messages)
}