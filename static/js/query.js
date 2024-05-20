inp = document.getElementById('input')
o = document.getElementById("output")
x = document.getElementById('button')
txt = document.getElementById("outputText")

function getQuery() {
    const data =  dbConnect(inp.value);
    putOutput(data)
}

function putOutput(data) {
    if (!data) {
        txt.innerText = undefined
    } else {
        txt.innerText = data
    }     
} 

x.addEventListener('click', getQuery)


