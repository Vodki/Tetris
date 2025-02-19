import { useState, useEffect, useCallback } from 'react';

const useWebSocket = (url: string) => {
  const [grid, setGrid] = useState<number[][]>([]);
  const [socket, setSocket] = useState<WebSocket | null>(null);

  // Initialize WebSocket connection
  useEffect(() => {
    const ws = new WebSocket(url);
    setSocket(ws);

    ws.onmessage = (event) => {
      try {
        console.log(event.data)
        const message: { type: string; data: string } = JSON.parse(event.data);
        
        if (message.type === 'GameUpdate') {
          setGrid(JSON.parse(message.data));
        }
      } catch (error) {
        console.error('Error handling message:', error);
      }
    };

    return () => {
      ws.close();
    };
  }, [url]);

  // Send message helper
  const sendMessage = useCallback((type: string, data: string) => {
    if (socket?.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify({ type, data }));
    }
  }, [socket]);

  return { grid, sendMessage };
};

export default useWebSocket;