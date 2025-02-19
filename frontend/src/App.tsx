import React from 'react';
import './App.css';
import WebSocketTest from './components/Tetris';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>Tetris</h1>
        < WebSocketTest/>
      </header>
    </div>
  );
}

export default App;
