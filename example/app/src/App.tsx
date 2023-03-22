import './App.css';
import { Outlet, Link } from "react-router-dom"

function App() {
  return (
    <div className="App">
      
      <Link to="/"><h1>Todo App</h1></Link>
      <Link to="/new">New Todo</Link>
      <Outlet />
    </div>
  );
}

export default App;
