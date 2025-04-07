package main

import (
    "html/template"
    "net/http"
)

type Message struct {
    User    string
    Content string
}

var messages []Message

func main() {
    http.HandleFunc("/", chatPage)
    http.HandleFunc("/send", handleSend)

    println("Servidor corriendo en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func chatPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("chat.html"))
    tmpl.Execute(w, messages)
}

func handleSend(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    userMsg := r.FormValue("message")
    messages = append(messages, Message{"Usuario", userMsg})

    // Respuesta básica del bot
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
