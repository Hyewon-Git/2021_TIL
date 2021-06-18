import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import Iframe from "react-iframe"

class Harbor extends Component {
  render() {
    return (
      <div className="Harbor">
          <p>Keycloak Login Success</p>
          <ul>
            <li><Link to="/secured">Main</Link></li>
            <li><Link to="/gitlab">Gitlab</Link></li>
            <li><Link to="/harbor">Harbor</Link></li>
            <li><Link to="/">Logout</Link></li>
          </ul>
        <center><p>HW- Harbor Page</p></center>
        <Iframe url="http://localhost:30002/c/oidc/login"
            width="100%"
            allow="geolocation *;"
            height="800px"
          />
      </div>
    );
  }
}
export default Harbor;