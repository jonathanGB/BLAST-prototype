const form = document.querySelector("form");
const output = document.getElementById("output");

form.addEventListener("submit", e => {
  e.preventDefault();
  
  const query = e.target[0].value;
  const isVerbose = e.target[1].checked;
  const kMer = parseInt(e.target[2].value);

  if (query.length < kMer) {
    alert("kmers can't be bigger than the query itself");
    return;
  }

  fetch("/blast", {
    method: "POST",
    body: JSON.stringify({
      query,
      isVerbose,
      kMer,
    })
  })
  .then(res => res.text())
  .then(text => output.innerText = text);
})