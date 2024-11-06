package internal

var userTasks = make(map[int][]string)

// AddTask добавляет задачу для пользователя
func AddTask(userID int, task string) {
    userTasks[userID] = append(userTasks[userID], task)
}

// ListTasks возвращает список задач пользователя
func ListTasks(userID int) []string {
    return userTasks[userID]
}

// RemoveTask удаляет задачу пользователя
func RemoveTask(userID int, task string) bool {
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

