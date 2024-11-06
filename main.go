package main

import (
    "log"
    "strings"
    "GriBotProdLev/internal"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    bot, err := tgbotapi.NewBotAPI(internal.BotToken)
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
            internal.AddTask(userID, task)
            response = "Задача добавлена: " + task

        case msgText == "/listtasks":
            tasks := internal.ListTasks(userID)
            if len(tasks) > 0 {
                response = "Ваши задачи:\n" + strings.Join(tasks, "\n")
            } else {
                response = "У вас нет активных задач."
            }

        case strings.HasPrefix(msgText, "/removetask "):
            task := strings.TrimSpace(strings.TrimPrefix(msgText, "/removetask "))
            if internal.RemoveTask(userID, task) {
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

