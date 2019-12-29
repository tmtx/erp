import React, { useState } from "react";
import { Pane, Button, TextInputField, toaster } from "evergreen-ui";
import API from "./../../Api";
import * as types from "./../../types";

const ProfileEditForm: React.FC = () => {
  const emptyMessages: types.ValidationMessages = {
    email: "",
  };
  const [validationMessages, setValidationMessages] = useState(emptyMessages);
  const [email, setEmail] = useState("");

  const submitForm = () => {
    API.post("/users/update", {email: email})
      .then( response => {
        if (response.data && response.data.status === "ok") {
          toaster.success("Successfully saved");
          setValidationMessages(emptyMessages);
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

  return (
    <Pane display="flex" padding={40} justifyContent="left" flexDirection="column">
      <h2>Change profile information</h2>
      <Pane width={300}>
        <TextInputField
          name="email"
          label="Email"
          onChange={ (e: React.ChangeEvent<HTMLInputElement>) => {
            if (!e) {
              return;
            }
            const target = e.target as HTMLInputElement;
            if (target) {
              setEmail(target.value)
            }
          }}
          placeholder="Email"
          isInvalid={getValidationMessage("email") !== null}
          validationMessage={getValidationMessage("email")}
        />
      </Pane>
      <Button iconBefore="floppy-disk" intent="success" width={100} onClick={submitForm}>
        Save
      </Button>
    </Pane>
  );
};

export default ProfileEditForm;
