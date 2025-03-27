import './App.css';
import Tetris from './components/Tetris';

function App() {

  return (
    <div className="App w-full h-full">
      <header className="App-header">
        <h1 className='text-4xl'>Welcome to my Tetris !</h1>
      </header>
        <h2>To play, please enter a Username</h2>
      < Tetris/>
    </div>
  );
}

export default App;
