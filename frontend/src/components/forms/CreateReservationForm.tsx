import React, { useState, useEffect } from "react";
import { Redirect } from "react-router-dom";
import {
  Pane,
  Button,
  TextInput,
  TextInputField,
  SelectField,
  FormField,
  toaster,
} from "evergreen-ui";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

import * as types from "./../../types";
import API from "./../../Api";

const CreateReservationForm: React.FC = () => {
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [spaceId, setSpaceId] = useState<number|null>(null);
  const [shouldRedirect, setShouldRedirect] = useState(false);

  const emptyMessages: types.ValidationMessages = {
    email: "",
    name: "",
    spaceId: "",
  };
  const [validationMessages, setValidationMessages] = useState(emptyMessages);
  const [availableRooms, setAvailableRooms] = useState<types.Space[]>([]);

  useEffect(() => {
    API.get("/spaces/available")
      .then(response => response.data)
      .then(data => {
        setAvailableRooms(data);
        setSpaceId(data[0].id);
      })
    ;

    // eslint-disable-next-line react-hooks/exhaustive-deps
    endDate.setDate(startDate.getDate() + 7);
  }, []);

  if (shouldRedirect) {
    return (
      <Redirect to="/" />
    );
  }

  const submitForm = () => {
    API.post("/reservations/create", {
      email: email,
      name: name,
      startDate: startDate,
      endDate: endDate,
      spaceId: spaceId,
    })
      .then( response => {
        if (response.data && response.data.status === "ok") {
          toaster.success("Successfully saved");
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

  return (
    <Pane
      display="flex"
      padding={40}
      justifyContent="left"
      flexDirection="column"
    >
      <h2>New reservation</h2>
      <Pane display="flex" justifyContent="space-between">
        <Pane width={300}>
          <TextInputField
            name="name"
            label="Name"
            placeholder="Name"
            onChange={ (e: React.ChangeEvent<HTMLInputElement>) => {
              if (!e) {
                return;
              }
              const target = e.target as HTMLInputElement;
              if (target) {
                setName(target.value)
              }
            }}
            isInvalid={getValidationMessage("name") !== null}
            validationMessage={getValidationMessage("name")}
          />
          <TextInputField
            name="email"
            label="Email"
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
        </Pane>
        <Pane width={600} paddingLeft={80}>
          <Pane display="flex" justifyContent="space-between" marginBottom="24px">
            <FormField label="Start date">
              <DatePicker
                selected={startDate}
                onChange={(date: Date) => setStartDate(date)}
                selectsStart
                startDate={startDate}
                endDate={endDate}
                customInput={<TextInput width={250} />}
              />
            </FormField>
            <FormField label="End date">
              <DatePicker
                selected={endDate}
                onChange={(date: Date) => setEndDate(date)}
                selectsEnd
                startDate={startDate}
                endDate={endDate}
                minDate={startDate}
                customInput={<TextInput width={250} />}
              />
            </FormField>
          </Pane>
          <Pane>
            <SelectField
              name="livingSpaceId"
              label="Room number"
              width={300}
              placeholder="Room number"
              onChange={ (e: React.ChangeEvent<HTMLInputElement>) => {
                if (!e) {
                  return;
                }
                const target = e.target as HTMLInputElement;
                if (target) {
                  setSpaceId(parseInt(target.value))
                }
              }}
            >
              { availableRooms.map((s: types.Space, i: number) => (
                <option key={i} value={s.id}>{s.id}</option>
              )) }
            </SelectField>
          </Pane>
        </Pane>
      </Pane>
      <div style={{marginTop: "20px"}} />
      <Button
        iconBefore="floppy-disk"
        intent="success"
        width={100}
        onClick={ () => submitForm() }
      >
        Save
      </Button>
    </Pane>
  );
};

export default CreateReservationForm;
