package main

var m = map[string]string{
	"commands/CreateTask":     "1",
	"commands/UpdateTask":     "2",
	"commands/DeleteTask":     "3",
	"commands/ProgressTask":   "4",
	"commands/CompleteTask":   "5",
	"commands/AssignTask":     "6",
	"commands/CryTask":        "7",
	"queries/QueryTask":       "8",
	"1commands/CreateTask":    "1",
	"1commands/UpdateTask":    "2",
	"1commands/DeleteTask":    "3",
	"1commands/ProgressTask":  "4",
	"1commands/CompleteTask":  "5",
	"1commands/AssignTask":    "6",
	"1commands/CryTask":       "7",
	"1queries/QueryTask":      "8",
	"2commands/CreateTask":    "1",
	"2commands/UpdateTask":    "2",
	"2commands/DeleteTask":    "3",
	"2commands/ProgressTask":  "4",
	"2commands/CompleteTask":  "5",
	"2commands/AssignTask":    "6",
	"2commands/CryTask":       "7",
	"2queries/QueryTask":      "8",
	"10commands/CreateTask":   "1",
	"10commands/UpdateTask":   "2",
	"10commands/DeleteTask":   "3",
	"10commands/ProgressTask": "4",
	"10commands/CompleteTask": "5",
	"10commands/AssignTask":   "6",
	"10commands/CryTask":      "7",
	"10queries/QueryTask":     "8",
	"11commands/CreateTask":   "1",
	"11commands/UpdateTask":   "2",
	"11commands/DeleteTask":   "3",
	"11commands/ProgressTask": "4",
	"11commands/CompleteTask": "5",
	"11commands/AssignTask":   "6",
	"11commands/CryTask":      "7",
	"11queries/QueryTask":     "8",
	"12commands/CreateTask":   "1",
	"12commands/UpdateTask":   "2",
	"12commands/DeleteTask":   "3",
	"12commands/ProgressTask": "4",
	"12commands/CompleteTask": "5",
	"12commands/AssignTask":   "6",
	"12commands/CryTask":      "7",
	"12queries/QueryTask":     "8",
}

func mapRoute(r string) string {
	return m[r]
}

func switchRoute(r string) string {
	switch r {
	case "commands/CreateTask":
		return "1"
	case "commands/UpdateTask":
		return "2"
	case "commands/DeleteTask":
		return "3"
	case "commands/ProgressTask":
		return "4"
	case "commands/CompleteTask":
		return "5"
	case "commands/AssignTask":
		return "6"
	case "commands/CryTask":
		return "7"
	case "queries/QueryTask":
		return "8"
	case "1commands/CreateTask":
		return "1"
	case "1commands/UpdateTask":
		return "2"
	case "1commands/DeleteTask":
		return "3"
	case "1commands/ProgressTask":
		return "4"
	case "1commands/CompleteTask":
		return "5"
	case "1commands/AssignTask":
		return "6"
	case "1commands/CryTask":
		return "7"
	case "1queries/QueryTask":
		return "8"
	case "2commands/CreateTask":
		return "1"
	case "2commands/UpdateTask":
		return "2"
	case "2commands/DeleteTask":
		return "3"
	case "2commands/ProgressTask":
		return "4"
	case "2commands/CompleteTask":
		return "5"
	case "2commands/AssignTask":
		return "6"
	case "2commands/CryTask":
		return "7"
	case "2queries/QueryTask":
		return "8"
	case "10commands/CreateTask":
		return "1"
	case "10commands/UpdateTask":
		return "2"
	case "10commands/DeleteTask":
		return "3"
	case "10commands/ProgressTask":
		return "4"
	case "10commands/CompleteTask":
		return "5"
	case "10commands/AssignTask":
		return "6"
	case "10commands/CryTask":
		return "7"
	case "10queries/QueryTask":
		return "8"
	case "11commands/CreateTask":
		return "1"
	case "11commands/UpdateTask":
		return "2"
	case "11commands/DeleteTask":
		return "3"
	case "11commands/ProgressTask":
		return "4"
	case "11commands/CompleteTask":
		return "5"
	case "11commands/AssignTask":
		return "6"
	case "11commands/CryTask":
		return "7"
	case "11queries/QueryTask":
		return "8"
	case "12commands/CreateTask":
		return "1"
	case "12commands/UpdateTask":
		return "2"
	case "12commands/DeleteTask":
		return "3"
	case "12commands/ProgressTask":
		return "4"
	case "12commands/CompleteTask":
		return "5"
	case "12commands/AssignTask":
		return "6"
	case "12commands/CryTask":
		return "7"
	case "12queries/QueryTask":
		return "8"
	}
	return ""
}