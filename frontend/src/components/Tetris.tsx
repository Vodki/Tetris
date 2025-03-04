import React, { useCallback, useEffect } from "react";
import useWebSocket from "../hooks/useWebSocket";
import { createEmptyGrid } from "../utils/gridUtils";

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
    sendMessage("game", "start");
  }, [sendMessage]);

  return (
    <div>
      level : {level}
      score : {score}
      <button onClick={handleStart}>Start Game</button>
      <Grid grid={currentGrid} />
    </div>
  );
};

export default Tetris;