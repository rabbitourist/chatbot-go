package main

import (
    "html/template"
    "net/http"
    "os"
)

type Message struct {
    User    string
    Content string
}

var messages []Message

func main() {
    http.HandleFunc("/", chatPage)
    http.HandleFunc("/send", handleSend)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // para pruebas locales
    }
    println("Servidor corriendo en el puerto " + port)
    http.ListenAndServe(":" + port, nil)
}

func chatPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("chat.html"))
    tmpl.Execute(w, messages)
}

func handleSend(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    userMsg := r.FormValue("message")
    messages = append(messages, Message{"Usuario", userMsg})

    var botResponse string
    switch userMsg {
    case "sí", "acepto":
        botResponse = "Perfecto. ¿Quieres explorar opción A o B?"
    case "opción A":
        botResponse = "Has elegido la opción A. ¿Quieres seguir por este camino?"
    default:
        botResponse = "¿Aceptarías esta idea: 'Explorar un futuro sostenible'?"
    }

    messages = append(messages, Message{"Bot", botResponse})
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
