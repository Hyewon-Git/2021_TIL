import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import Iframe from "react-iframe"

class Welcome extends Component {
  render() {
    return (
      <div className="Welcome">
        <p>Welcome HW Page</p>
        <Link to="/secured">Keycloak Login</Link>
      </div>
    );
  }
}
export default Welcome;