import React from 'react';
import Sidebar from './components/Sidebar';
import Content from './components/Content';

function App() {
  return (
    <div className="flex">
      <Sidebar />
      <Content />
    </div>
  );
}

export default App;
