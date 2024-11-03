package main

import (
    "log"
    "strings"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Временное хранилище задач
var userTasks = make(map[int][]string)

func main() {
    bot, err := tgbotapi.NewBotAPI("7688165433:AAFSyeoSs80V8DJGRF_X0wBqyzlAqwwKlx4")
    if err != nil {
        log.Panic(err)
    }
    bot.Debug = true

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates, _ := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil { // игнорируем не-текстовые обновления
            continue
        }

        userID := update.Message.From.ID
        msgText := strings.ToLower(update.Message.Text)
        response := ""

        switch {
        case strings.HasPrefix(msgText, "/addtask "):
            task := strings.TrimSpace(strings.TrimPrefix(msgText, "/addtask "))
            addTask(userID, task)
            response = "Задача добавлена: " + task

        case msgText == "/listtasks":
            tasks := listTasks(userID)
            if len(tasks) > 0 {
                response = "Ваши задачи:\n" + strings.Join(tasks, "\n")
            } else {
                response = "У вас нет активных задач."
            }

        case strings.HasPrefix(msgText, "/removetask "):
            task := strings.TrimSpace(strings.TrimPrefix(msgText, "/removetask "))
            if removeTask(userID, task) {
                response = "Задача удалена: " + task
            } else {
                response = "Задача не найдена."
            }

        default:
            response = "Доступные команды:\n" +
                "/addtask <задача> — добавить задачу\n" +
                "/listtasks — показать все задачи\n" +
                "/removetask <задача> — удалить задачу"
        }

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
        bot.Send(msg)
    }
}

func addTask(userID int, task string) {
    userTasks[userID] = append(userTasks[userID], task)
}

func listTasks(userID int) []string {
    return userTasks[userID]
}

func removeTask(userID int, task string) bool {
    tasks, exists := userTasks[userID]
    if !exists {
        return false
    }

    for i, t := range tasks {
        if t == task {
            userTasks[userID] = append(tasks[:i], tasks[i+1:]...)
            return true
        }
    }
    return false
}

