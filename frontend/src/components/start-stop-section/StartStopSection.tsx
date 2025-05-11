import React, { useEffect, useState } from "react";
import { useWebSocket } from "../../context/WebSocketContext";

type InterfaceItem = {
  name: string;
  description: string;
};

function StartStopSection() {
  const { send, isConnected, addMessageListener } = useWebSocket();
  const [interfaces, setInterfaces] = useState<InterfaceItem[]>([]);

  useEffect(() => {
    if (isConnected) {
      send({ action: "get-iface" });
    }
  }, [isConnected]);

  useEffect(() => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const unsubscribe = addMessageListener((msg: any) => {
      console.log("📥 Received in StartStopSection:", msg);

      if (msg.type === "iface-list" && Array.isArray(msg.data)) {
        setInterfaces(msg.data);
      }
    });

    return () => unsubscribe(); // Clean up on unmount
  }, []);

  return (
    <div className="start-stop-section">
      <h1>Choose Interface</h1>
      <select name="interfaces" id="interfaces">
        <option value="none" selected>
          Select Interface
        </option>
        {interfaces.map((iface, index) => (
          <option key={index} value={iface.name}>
            {iface.description || iface.name}
          </option>
        ))}
      </select>
      <button>Start</button>
    </div>
  );
}

export default StartStopSection;
