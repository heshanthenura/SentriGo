const ifaceSelect = document.getElementById("interface");
const serverStatus = document.querySelector(".server-status");
const startStopBtn = document.getElementById("start-btn");

const ws = new WebSocket("ws://localhost:8080/ws");

ws.onopen = function () {
  console.log("Connected to WebSocket");

  serverStatus.classList.remove("offline");
  serverStatus.classList.add("online");
  serverStatus.textContent = "ONLINE";
};

ws.onclose = function () {
  console.log("Disconnected from WebSocket");
  serverStatus.classList.remove("online");
  serverStatus.classList.add("offline");
  serverStatus.textContent = "OFFLINE";
};

ws.onerror = function (error) {
  console.error("WebSocket error:", error);
  serverStatus.classList.remove("online");
  serverStatus.classList.add("offline");
  serverStatus.textContent = "OFFLINE";
};

ifaceSelect.addEventListener("change", () => {
  ws.send(
    JSON.stringify({
      action: "toggle-iface",
      iface: ifaceSelect.value,
    })
  );
});

startStopBtn.addEventListener("click", () => {
  if (ifaceSelect.value !== "none") {
    if (startStopBtn.textContent === "STOP") {
      startStopBtn.textContent = "START";
      startStopBtn.classList.remove("STOP");
      startStopBtn.classList.add("START");
    } else {
      startStopBtn.textContent = "STOP";
      startStopBtn.classList.remove("START");
      startStopBtn.classList.add("STOP");
    }
    console.log("log " + ifaceSelect.value);

    ws.send(
      JSON.stringify({
        action: "start-stop",
      })
    );
  } else {
    alert("Please select a valid interface before starting.");
  }
});
