import React, { useState, useEffect } from "react";
import { Pane } from "evergreen-ui";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";

import LogInForm from "./components/forms/LogInForm";
import ReservationList from "./components/ReservationList";
import Header from "./components/Header";
import CreateReservationForm from "./components/forms/CreateReservationForm";
import Api from "./Api";
import { UserContext } from "./userContext";

import * as types from "./types";

const App: React.FC = () => {
  const [currentUser, setCurrentUser] = useState<types.User|null>(null);

  const getSessionData = () => {
    Api.get("/users/me")
      .then( response => {
        if (response.data && response.data.email && response.data.id) {
          setCurrentUser({
            email: response.data.email,
            id: response.data.id
          });
        }
      });
  };

  useEffect(() => {
    getSessionData();
  }, []);

  const privateRoutes = () => {
    if (!currentUser) {
      return (
        <Redirect to="/login" />
      );
    }

    return (
      <Switch>
        <Route exact path="/create-reservation">
          <CreateReservationForm />
        </Route>
        <Route exact path="/">
          <ReservationList />
        </Route>
      </Switch>
    );
  };

  const loginBox = () => {
    if (currentUser) {
      return (
        <Redirect to="/" />
      );
    }

    return (
      <Pane
        clearfix
        justifyContent="center"
        alignItems="center"
        width="100%"
        height="100%"
        display="flex"
        flexDirection="column"
        position="relative"
      >
        <LogInForm getSessionData={getSessionData} />
      </Pane>
    );
  }

  return (
    <UserContext.Provider value={currentUser}>
      <div className="App">
        <Router basename={process.env.REACT_APP_BASEURL}>
          { currentUser ?
              <Header />
              :
              null
          }
          <Pane
              display="flex"
              marginLeft="auto"
              marginRight="auto"
              width="1000px"
              justifyContent="center"
              position="relative"
              flexDirection="column"
          >
            <Switch>
              <Route path="/login">
                { loginBox() }
              </Route>
              { privateRoutes() }
            </Switch>
          </Pane>
        </Router>
      </div>
    </UserContext.Provider>
  );
}

export default App;
