import Tetris from "@/components/Tetris";

export function GamePage({ username }: { username: string }) {
  return (
    <div className="game-page">
      <h2>Welcome, {username}!</h2>
      <Tetris />
    </div>
  );
}