import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import Iframe from "react-iframe"

class Gitlab extends Component {
  render() {
    return (
      <div className="Gitlab">
          <p>Keycloak Login Success</p>
          <ul>
            <li><Link to="/secured">Main</Link></li>
            <li><Link to="/gitlab">Gitlab</Link></li>
            <li><Link to="/harbor">Harbor</Link></li>
            <li><Link to="/">Logout</Link></li>
          </ul>
        <center><p>HW- Gitlab Page</p></center>
        <Iframe url="https://hw2.dudaji.com:8443/auth/realms/dudaji-100/protocol/saml/clients/gitlab"
            width="100%"
            allow="geolocation *;"
            height="800px"
          />
      </div>
    );
  }
}
export default Gitlab;