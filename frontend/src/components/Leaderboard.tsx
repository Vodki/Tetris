import React from "react";

interface LeaderboardEntry {
  username: string;
  score: number;
}

interface LeaderboardProps {
  entries: LeaderboardEntry[];
}

const Leaderboard: React.FC<LeaderboardProps> = ({ entries }) => {
  return (
    <div className="leaderboard-container">
      <h2 className="leaderboard-title">Leaderboard</h2>
      <div className="leaderboard-list">
        {entries.map((entry, index) => (
          <div key={`${entry.username}-${index}`} className="leaderboard-item">
            <div className="leaderboard-rank">
              <span className="rank-number">{index + 1}.</span>
              <span className="username">{entry.username}</span>
            </div>
            <span className="score">{entry.score}</span>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Leaderboard;