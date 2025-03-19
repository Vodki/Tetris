import './App.css';
import Tetris from './components/Tetris';
import { Input } from "@/components/ui/input"
import { Button } from './components/ui/button';
import { useState } from 'react';
import { uniqueNamesGenerator, adjectives, colors, animals } from "unique-names-generator";


function App() {

  const [username, setUsername] = useState<string>("")
  const handleGeneratePseudo = () => {
    const generatedName = uniqueNamesGenerator({
      dictionaries: [adjectives, colors, animals],
      separator: "_",
      style: "capital",
      length: 3
    }) + Math.floor(Math.random() * 90 + 10);

    setUsername(generatedName);
  };

  return (
    <div className="App w-full h-full">
      <header className="App-header">
        <h1 className='bg-red-500'>Welcome to my Tetris !</h1>
      </header>
        <h2>To play, please enter a Username</h2>
      < Tetris/>
    </div>
  );
}

export default App;
