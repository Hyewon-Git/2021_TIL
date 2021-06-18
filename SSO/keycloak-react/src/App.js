import React, { Component } from 'react';
import { BrowserRouter, Route, Link } from 'react-router-dom';
import Welcome from './Welcome';
import Secured from './Secured';
import Gitlab from './Gitlab';
import Harbor from './Harbor';

import './App.css';

class App extends Component {

  render() {
    return (
      <BrowserRouter>
        <div className="container">
          <center><Route exact path="/" component={Welcome} /></center>
          {/* <ul>
            <li><Link to="/">Main</Link></li>
          </ul> */}
          <Route path="/gitlab" component={Gitlab} />
          <Route path="/harbor" component={Harbor} />
          <Route path="/secured" component={Secured} />
          {/* <Route  path="/gitlab" render={() => (window.location ="http://hw2.dudaji.com:30010")} />
          <Route  path="/harbor" render={() => (window.location ="http://localhost:30002")} /> */}
        </div>
      </BrowserRouter>
    );
  }
}
export default App;