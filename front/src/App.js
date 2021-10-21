import React from 'react';
import './App.css';

import Search from  './components/Search'
import Nav from  './components/Nav'


function App() {
  return (
    <div className="App">
    <div className="App-header">
        <Nav />
     <Search />
     </div>
    </div>
  );
}

export default App;
