import React from "react";
import { Pane } from "evergreen-ui";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";

import LogInBox from "./components/LogInBox";
import ReservationList from "./components/ReservationList";
import Header from "./components/Header";
import CreateReservationForm from "./components/CreateReservationForm";

interface User {
  email: string|null;
}

const App: React.FC = () => {

  const getCurrentUser = () => {
    // TODO: get current user session data from /users/current
    return null;
  }

  const privateRoutes = () => {
    if (!getCurrentUser()) {
      return (
        <Redirect to="/login" />
      );
    }

    return (
      <div>
        <Route path="/create-reservation">
          <CreateReservationForm />
        </Route>
        <Route path="/">
          <ReservationList />
        </Route>
      </div>
    );
  };

  return (
    <div className="App">
      <Router>
        <Header />
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
                <LogInBox />
              </Pane>
            </Route>
            { privateRoutes() }
          </Switch>
        </Pane>
      </Router>
    </div>
  );
}

export default App;
