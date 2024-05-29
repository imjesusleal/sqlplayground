const playground = document.createElement("div")
const textarea = document.createElement("textarea")
const button = document.createElement("button")
const output = document.createElement("output")
const parraf = document.createElement("div")

function renderPlayground(before_node) {
    playground.id = "playground"
    playground.classList = "grid place-items-center py-8 gap-4"

    textarea.classList = "textarea textarea-success text-white textarea-md w-2/4"
    textarea.placeholder = "create your query"
    textarea.id = "input"

    button.class = "btn"
    button.id = "button"
    button.innerHTML = "Click Me"

    output.id="output"

    parraf.id = "outputText"

    playground.appendChild(textarea)
    playground.appendChild(button)
    playground.appendChild(output)
    output.appendChild(parraf)
    x = document.getElementById(before_node)
    x.after(playground)
}

renderPlayground("seg1")

