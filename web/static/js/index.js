document.getElementById("interface").addEventListener("change", function () {
  const selectedValue = this.value;
  console.log("You selected:", selectedValue);
  const formData = new FormData();
  formData.append("iface", selectedValue);
  const xhttp = new XMLHttpRequest();
  xhttp.open("POST", "/interface", true);
  xhttp.onreadystatechange = function () {
    if (this.readyState === 4 && this.status === 200) {
      console.log("Response:", this.responseText);
    }
  };

  xhttp.send(formData);
});

document.getElementById("start-btn").addEventListener("click", function () {
  alert();
});
