inp = document.getElementById('input')
o = document.getElementById("output")
x = document.getElementById('button')
txt = document.getElementById("outputText")

function getQuery() {
    const data =  dbConnect(inp.value);
    putOutput(data)
}

function putOutput(data) {
    txt.innerText = ""
    if (!data) {
        txt.innerText = undefined
    } else {
        switch (typeof(data)) {
            case "string":
            txt.innerText += data
            break;
            case "object":
            mapData(data)
            break;
            default:
            txt.innerText += data
            break;
        }
    }     
} 

function mapData(data) {
    for (i=0;i<data.length; i++) {
        Object.entries(data[i]).map(e => {
            let keys = e[0]
            let val = e[1]
            txt.innerText += "\n" + keys + "\t" + val
            txt.scrollIntoView()
        })
    }
}

x.addEventListener('click', getQuery)


