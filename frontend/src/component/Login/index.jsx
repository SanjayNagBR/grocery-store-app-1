import React, { Component } from "react";
import { Link, Redirect } from "react-router-dom";
import axios from "axios";
import "./index.css";

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: "",
      password: "",
      userLogged: false,
      invalidCreds: false
    };
  }

  emailChangeHandler = e => {
    this.setState({
      email: e.target.value
    });
  };

  passwordChangeHandler = e => {
    this.setState({
      password: e.target.value
    });
  };

  userLogin = e => {
    e.preventDefault();

    axios
      .get(`http://localhost:3001/users/${this.state.email}`)
      .then(response => {
        console.log("Status Code : ", response.status);
        if (response.status === 200) {
          console.log("response data:", response);
          if (response.data.password === this.state.password) {
            sessionStorage.setItem("email", response.data.email);
            this.setState({
              userLogged: true,
              invalidCreds: false
            });
          } else {
            this.setState({
              userLogged: false,
              invalidCreds: true
            });
          }
        }
      })
      .catch(err => {
        this.setState({
          userLogged: false,
          invalidCreds: true
        });
      });
  };

  render() {
    if (this.state.userLogged) {
      return <Redirect to="/inventory" />;
    }
    return (
      <React.Fragment>
        <div id="login-page-headline" className="container">
          <h2>Log in to Shayona Grocery Store</h2>
          <br />
          <h4>
            Need an account? <Link to={{ pathname: "/signup" }}>Sign Up</Link>
          </h4>
          <h6 style={{ color: "#BA160C" }}>
            {this.state.invalidCreds &&
              "The email and password you entered did not match our records. Please double-check and try again."}
          </h6>
          <div className="login-form card rounded-0">
            <div className="card-header card-title">
              <h3>User Account Login</h3>
            </div>
            <div className="card-body">
              <form method="POST">
                <div className="form-group">
                  <input
                    onChange={this.emailChangeHandler}
                    type="text"
                    className="form-control rounded-0"
                    placeholder="Email address"
                  />
                </div>
                <div className="form-group">
                  <input
                    onChange={this.passwordChangeHandler}
                    type="password"
                    className="form-control rounded-0"
                    placeholder="Password"
                  />
                </div>
                <h6>Forgot password?</h6>
                <button
                  onClick={this.userLogin}
                  className="btn btn-block btn-login rounded-0"
                >
                  Log In
                </button>
              </form>
            </div>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

export default Login;
