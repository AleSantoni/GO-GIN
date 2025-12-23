function changeMessage(){
    const messageElement = document.getElementById('message');
    const currentMessage = messageElement.textContent;
    const alternateMessage = [
        "Hello de Nuevo", 
        "Welcome Again",
        "Que tenga un Buen Dia",
        "Hola de Nuevo",
        "Que tenga un Buen Dia"];

        let newMessage;
        do {
            newMessage = alternateMessage[Math.floor(Math.random() * alternateMessage.length)];
        } while (newMessage === currentMessage);

        messageElement.textContent = newMessage;
}