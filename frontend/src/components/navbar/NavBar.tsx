import React from "react";
import { Link } from "react-router-dom";
import "./NavBar.css";
import { useWebSocket } from "../../context/WebSocketContext";

function NavBar() {
  const { isConnected } = useWebSocket();

  return (
    <nav className="navbar">
      <div className="logo">SentriGO</div>
      <div className="links">
        <Link to="/">Home</Link>
        <Link to="/about">About</Link>
        <p>
          WebSocket status: {isConnected ? "✅ Connected" : "❌ Disconnected"}
        </p>
      </div>
    </nav>
  );
}

export default NavBar;
