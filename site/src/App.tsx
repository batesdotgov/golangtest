import React from "react";
import "./App.css";
import { BrowserRouter, NavLink, Switch, Route } from "react-router-dom";

function App() {
  return (
    <div className='App'>
      <BrowserRouter>
        <NavLink to='/'>Home</NavLink>
        <NavLink to='/about'>about</NavLink>
        <Switch>
          <Route path='/' exact render={() => "home"}></Route>
          <Route path='/about' render={() => "about"}></Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
