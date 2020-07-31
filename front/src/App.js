import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          电影饭-人工智能电影推荐系统
        </p>
        <a
          className="App-link"
          href="https://www.80shihua.com/"
          target="_blank"
          rel="noopener noreferrer"
        >
          by梦回故里
        </a>
      </header>
    </div>
  );
}

export default App;
