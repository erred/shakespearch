const Controller = {
  search: (ev) => {
    ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    const response = fetch(`/search?q=${data.query}`).then((response) => {
      response.json().then((results) => {
        Controller.updateTable(results);
      });
    });
  },

  updateTable: (results) => {
    const table = document.getElementById("table-body");
    const rows = [];
    rows.push(`<th><td>Play</td><td>Speaker</td><td>Text</td></th>`);
    for (let result of results) {
      rows.push(`<tr><td>${result.Play}</td><td>${result.Speaker}</td><td>${result.Text}</td></tr>`);
    }
    table.innerHTML = rows;
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
