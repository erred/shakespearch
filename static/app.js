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
    rows.push(`<thead><th>Play</th><th>Speaker</th><th>Text</th></thead>`);
    for (let result of results) {
      rows.push(`<tr><td>${result.Play}</td><td>${result.Speaker}</td><td>${result.Text}</td></tr>`);
    }
    table.innerHTML = rows.join("\n");
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
