import React, { useCallback, useEffect, useState } from "react";
import useWebSocket from "../hooks/useWebSocket";
import { createEmptyGrid } from "../utils/gridUtils";
import { Button } from "./ui/button";
import { uniqueNamesGenerator, adjectives, colors, animals } from "unique-names-generator";
import { Input } from "./ui/input";

interface ApiResponse {
  message: string;
}

const Grid: React.FC<{ grid: number[][] }> = React.memo(({ grid }) => (
  <div className="grid">
    {grid.map((row, rowIndex) => (
      <div key={rowIndex} className="row">
        {row.map((cell, cellIndex) => (
          <div
            key={cellIndex}
            className={`cell ${cell !== 0 ? `color-${cell}` : ''}`}
          />
        ))}
      </div>
    ))}
  </div>
));


const Tetris: React.FC = () => {
  const defaultGrid = createEmptyGrid();
  const { grid, sendMessage, score, level, gameOn } = useWebSocket("ws://localhost:8080/ws");
  const [username, setUsername] = useState<string>("")
  const [data, setData] = useState<ApiResponse | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  
  const handleGeneratePseudo = () => {
    const generatedName = uniqueNamesGenerator({
      dictionaries: [adjectives, colors, animals],
      separator: "_",
      style: "capital",
      length: 3
    }) + Math.floor(Math.random() * 90 + 10);
  
    setUsername(generatedName);
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/leaderboard');

        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        } else {
          console.log("response = ", response)
        }

        const result: ApiResponse = await response.json();
        setData(result);
        setError(null);
      } catch (error) {
        console.log("CATCHED ERROR")
        if (error instanceof Error) {
          setError(error.message);
          console.log(error.message)
        } else {
          setError('An unknown error occurred');
        }
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  // Always show current grid or default if empty
  const currentGrid = grid && grid.length > 0 ? grid : defaultGrid;

  const handleKeyDown = useCallback((event: KeyboardEvent) => {
    if (!gameOn) return;
    switch (event.key) {
      case "ArrowUp":
        sendMessage("game", "Rotate");
        event.preventDefault();
        break;
      case "ArrowRight":
        sendMessage("game", "MoveRight");
        event.preventDefault();
        break;
      case "ArrowLeft":
        sendMessage("game", "MoveLeft");
        event.preventDefault();
        break;
      case "ArrowDown":
        sendMessage("game", "MoveDown");
        event.preventDefault();
        break;
      case " ":
        sendMessage("game", "HardDrop");
        event.preventDefault();
        break;
      default:
        break;
    }
  }, [sendMessage, gameOn]);

  // Add keyboard event listener
  useEffect(() => {
    document.addEventListener("keydown", handleKeyDown);
    return () => {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, [handleKeyDown]);

  const handleStart = useCallback(() => {
    if (!gameOn)
      sendMessage("start", username);
  }, [sendMessage, gameOn, username]);

  return (
    <div style={{
      display: "flex",
      flexDirection: "row",
      justifyContent: "space-around",
      alignItems: "center",
      columnGap: "3rem"
    }}>
      <div>
        <Input
          placeholder='Username'
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          />
        <Button
          onClick={handleGeneratePseudo}>
          Create a random Username
        </Button>
        <p>level : {level}</p>
        <p>score : {score}</p>
        <Button
          className='disabled:opacity-50 disabled:cursor-not-allowed'
          disabled={!username.trim() || gameOn}
          onClick={handleStart}>
          Start Game / Replay
        </Button>
      </div>
      <Grid grid={currentGrid} />
    </div>
  );
};

export default Tetris;