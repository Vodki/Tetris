import { useState, useEffect, useCallback } from 'react';

const useWebSocket = (url: string) => {
  const [grid, setGrid] = useState<number[][]>([]);
  const [score, setScore] = useState<number>();
  const [gameOn, setGameOn] = useState<boolean>(false);
  const [level, setLevel] = useState<number>();
  const [socket, setSocket] = useState<WebSocket | null>(null);

  // Initialize WebSocket connection
  useEffect(() => {
    const ws = new WebSocket(url);
    setSocket(ws);

    ws.onmessage = (event) => {
      try {
        console.log(event.data)
        const message: { type: string; data: string; score: number; level:number; gameOn: boolean } = JSON.parse(event.data);
        
        if (message.type === 'GameUpdate') {
          setGrid(JSON.parse(message.data));
          setScore(message.score)
          setGameOn(message.gameOn)
          setLevel(message.level)
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

  return { grid, sendMessage, score, level, gameOn };
};

export default useWebSocket;