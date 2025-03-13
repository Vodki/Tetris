import React, { useCallback, useEffect } from "react";
import useWebSocket from "../hooks/useWebSocket";
import { createEmptyGrid } from "../utils/gridUtils";
import { Button } from "./ui/button";

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
      sendMessage("game", "start");
  }, [sendMessage, gameOn]);

  return (
    <div style={{
      display: "flex",
      flexDirection: "row",
      justifyContent: "space-around",
      alignItems: "center",
      columnGap: "3rem"
    }}>
      <div>
        <p>level : {level}</p>
        <p>score : {score}</p>
        <Button onClick={handleStart}>Start Game / Replay</Button>
      </div>
      <Grid grid={currentGrid} />
    </div>
  );
};

export default Tetris;