var inp = document.getElementById('input')
var o = document.getElementById("output")
var x = document.getElementById('button')
var txt = document.getElementById("outputText")

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
            generateTable(data)
            break;
            default:
            txt.innerText += data
            break;
        }
    }     
} 

function generateTable(obj) {
    const keys = Object.keys(obj);
    if (keys.length === 0) return '';

    const columns = Object.keys(obj[keys[0]]);

    let table = '<table class="table table-md"><thead><tr>';
    columns.forEach(columna => {
        table += `<th class="text-base">${columna}</th>`;
    });
    table += '</tr></thead><tbody>';

    keys.forEach(key => {
        table += '<tr>';
        columns.forEach(column => {
            table += `<td class="text-base">${obj[key][column]}</td>`;
        });
        table += '</tr>';
    });

    table += '</tbody></table>';
    console.log(table)
    txt.innerHTML = table
    txt.scrollIntoView()
}

x.addEventListener('click', getQuery)


