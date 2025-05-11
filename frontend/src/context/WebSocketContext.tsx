import React, {
  createContext,
  useContext,
  useEffect,
  useRef,
  useState,
} from "react";

// Define the WebSocket context type
type WebSocketContextType = {
  send: (data: unknown) => void;
  isConnected: boolean;
  addMessageListener: (cb: (msg: unknown) => void) => () => void;
};

// Create the context
const WebSocketContext = createContext<WebSocketContextType | null>(null);

// WebSocket Provider Component
export const WebSocketProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const socket = useRef<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState(false);

  // List of message listeners (subscribers)
  const listeners = useRef<((msg: unknown) => void)[]>([]);

  // Add message listener and return unsubscribe function
  const addMessageListener = (cb: (msg: unknown) => void) => {
    listeners.current.push(cb);
    return () => {
      listeners.current = listeners.current.filter(
        (listener) => listener !== cb
      );
    };
  };

  // Open WebSocket connection on mount
  useEffect(() => {
    const wsUrl =
      window.location.protocol === "https:"
        ? `wss://${window.location.host}/ws`
        : `ws://${window.location.hostname}:8080/ws`;

    socket.current = new WebSocket(wsUrl);

    socket.current.onopen = () => {
      console.log("✅ WebSocket connected");
      setIsConnected(true);
    };

    socket.current.onclose = () => {
      console.log("🔌 WebSocket disconnected");
      setIsConnected(false);
    };

    socket.current.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        console.log("📨 Received:", data);
        listeners.current.forEach((cb) => cb(data));
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
      } catch (e) {
        console.warn("⚠️ Failed to parse WebSocket message:", event.data);
      }
    };

    return () => {
      socket.current?.close();
    };
  }, []);

  // Send data to WebSocket server
  const send = (data: unknown) => {
    if (socket.current?.readyState === WebSocket.OPEN) {
      socket.current.send(JSON.stringify(data));
    } else {
      console.warn("⚠️ WebSocket not open. Could not send message.");
    }
  };

  // Provide context values
  return (
    <WebSocketContext.Provider
      value={{ send, isConnected, addMessageListener }}
    >
      {children}
    </WebSocketContext.Provider>
  );
};

// Custom hook to use the WebSocket context
export const useWebSocket = () => {
  const context = useContext(WebSocketContext);
  if (!context) {
    throw new Error("useWebSocket must be used within a WebSocketProvider");
  }
  return context;
};
