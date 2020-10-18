const socket = io({transports: ['websocket']});

// listen for new clipboard
socket.on('new-copy', function(message) {
    const elem = document.getElementById("field")
    elem.value = message
});

// Listen for new connection amount
socket.on("new-count", function (message){
    let elem = document.getElementById("count")
    elem.innerText = message
    elem.classList.remove("counter-anim")
    void elem.offsetWidth
    elem.classList.add("counter-anim")
})


// Prevent reloading when accidentally submitting form.
document.getElementById("form").addEventListener("submit", function(event){
    event.preventDefault()
});

// Update clipboard on input.
document.getElementById("field").addEventListener("input", function(event) {
        if (event.target && event.target.value){
            socket.emit("update", event.target.value)
        }
});

// Update clipboard on delete and backspace
document.getElementById("field").addEventListener("keydown", function (event){
    if (event.code === 46 || event.code === 8){
        socket.emit("update", event.target.value)
    }
});

document.getElementById("field").addEventListener("change", function (event){
        socket.emit("update", event.target.value)
});

// Copy to clipboard
document.getElementById("copy").addEventListener("click", function (event){
    const field = document.getElementById("field")
    field.select()
    field.setSelectionRange(0,99999)
    document.execCommand("copy")
});

// Paste from clipboard
document.getElementById("paste").addEventListener("click", async function (event){
    const text = await navigator.clipboard.readText();
    const field = document.getElementById("field")
    field.value = text
    field.focus()
    socket.emit("update", text)
})
