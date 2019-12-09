import React, { useState, useEffect } from "react";
import { Pane, Button, TextInput, TextInputField, SelectField, FormField } from "evergreen-ui";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

import API from './../Api';

interface Space {
  id: number;
}

const submitForm = () => {
  // TODO: implement email change, validation, etc.
};

const CreateReservationForm: React.FC = () => {
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [availableRooms, setAvailableRooms] = useState<Space[]>([]);
  endDate.setDate(startDate.getDate() + 7);

  useEffect(() => {
    API.get("/spaces/available")
      .then(response => response.data)
      .then(data => setAvailableRooms(data))
    ;
  }, []);

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
          />
          <TextInputField
            name="email"
            label="Email"
            placeholder="Email"
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
            >
              { availableRooms.map((s: Space) => (
                <option key="{s.id}" value="{s.id}">{s.id}</option>
              )) }
            </SelectField>
          </Pane>
        </Pane>
      </Pane>
      <div style={{marginTop: "20px"}} />
      <Button iconBefore="floppy-disk" intent="success" width={100}>
        Save
      </Button>
    </Pane>
  );
};

export default CreateReservationForm;
