
import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import Keycloak from 'keycloak-js'
import UserInfo from './UserInfo';
import Logout from './Logout';

class Secured extends Component {

  constructor(props) {
    super(props);
    this.state = { keycloak: null, authenticated: false };
  }
  
  componentDidMount() {
    const keycloak = Keycloak('/keycloak.json');
    keycloak.init({onLoad: 'login-required'}).then(authenticated => {
      this.setState({ keycloak: keycloak, authenticated: authenticated })
    })
  }

  render() {
    if (this.state.keycloak) {
      if (this.state.authenticated) return (
        <div>
          <p>Keycloak Login Success</p>
          <ul>
            <li><Link to="/secured">Main</Link></li>
            <li><Link to="/gitlab">Gitlab</Link></li>
            <li><Link to="/harbor">Harbor</Link></li>
            <li><Link to="/">Logout</Link></li>
          </ul>
          <UserInfo keycloak={this.state.keycloak} />

          {/* <Logout keycloak={this.state.keycloak} /> */}

          {/* hw-dudaji.com -> localhost:3000
          hw-dudaji.com/gitlab -> localhost:30010
          hw-dudaji.com/harbor -> localhost:30002 */}
          
        </div>
      ); else return (<div>Unable to authenticate!</div>)
    }
    return (
      <div>Initializing Keycloak...</div>
    );
  }
}
export default Secured;