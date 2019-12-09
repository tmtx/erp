import React, { useState } from "react";
import { Pane, Button, Heading, TextInputField } from "evergreen-ui";
import { Redirect } from "react-router-dom";
import API from "./../Api";

// TODO: DRY - see ProfileEditForm
interface ValidationMessages {
  [index: string]: string;
};

const LogInBox: React.FC = () => {
  const emptyMessages: ValidationMessages = {
    email: "",
    password: "",
  };
  const [validationMessages, setValidationMessages] = useState(emptyMessages);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [shouldRedirect, setShouldRedirect] = useState(false);

  const submitForm = () => {
    API.post("/users/login", {email: email, password: password})
      .then( response => {
        if (response.data && response.data.status === "ok") {
          setValidationMessages(emptyMessages);
          setShouldRedirect(true);
        } else if (response.data && response.data.status === "error") {
          setValidationMessages(response.data.errors);
        }
      });
  };

  const getValidationMessage = (key: string): string|null => {
    if (!validationMessages[key] || validationMessages[key].length === 0) {
      return null;
    }

    return validationMessages[key];
  };

  if (shouldRedirect) {
    return (
      <Redirect to="/" />
    );
  }

  return (
    <Pane marginTop="15%" display="flex" alignItems="center" flexDirection="column">
      <Heading size={800}>Authenticate</Heading>
      <Pane
        elevation={3}
        backgroundColor="white"
        width={300}
        height={220}
        margin={24}
        display="flex"
        justifyContent="center"
        alignItems="center"
        flexDirection="column"
      >
        <TextInputField
          marginTop={15}
          marginBottom={20}
          height={35}
          label=""
          name="email"
          placeholder="Email"
          onChange={ (e: React.ChangeEvent<HTMLInputElement>) => {
            if (!e) {
              return;
            }
            const target = e.target as HTMLInputElement;
            if (target) {
              setEmail(target.value)
            }
          }}
          isInvalid={getValidationMessage("email") !== null}
          validationMessage={getValidationMessage("email")}
        />
        <TextInputField
          marginTop={10}
          marginBottom={30}
          height={35}
          label=""
          name="password"
          type="password"
          placeholder="Password"
          onChange={ (e: React.ChangeEvent<HTMLInputElement>) => {
            if (!e) {
              return;
            }
            const target = e.target as HTMLInputElement;
            if (target) {
              setPassword(target.value)
            }
          }}
          isInvalid={getValidationMessage("password") !== null}
          validationMessage={getValidationMessage("password")}
        />
        <Button onClick={ () => submitForm() }>Log in</Button>
      </Pane>
    </Pane>
  );
}

export default LogInBox;
