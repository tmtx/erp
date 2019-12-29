import React, { useState, useEffect } from "react";
import { Table, Pane, Button } from "evergreen-ui";
import { Link } from "react-router-dom";
import * as types from "./../types";
import API from "./../Api";

const ReservationList: React.FC = () => {
  const [reservations, setReservations] = useState<types.Reservation[]>([]);

  useEffect(() => {
    API.get("/reservations/list")
      .then(response => response.data)
      .then(data => {
        setReservations(data);
      })
    ;
  }, []);

  return (
    <Pane width="100%">
      <Pane
        height={120}
        display="flex"
        marginTop="auto"
        marginBottom="auto"
        alignItems="center"
      >
        <Button
          height={30}
          marginRight={16}
          appearance="primary"
          intent="success"
          iconBefore="add"
          is={Link}
          to="/create-reservation"
        >
          Create reservation
        </Button>
      </Pane>
      <Table>
        <Table.Head>
          <Table.TextHeaderCell>Name</Table.TextHeaderCell>
          <Table.TextHeaderCell>Email</Table.TextHeaderCell>
          <Table.TextHeaderCell>Room number</Table.TextHeaderCell>
          <Table.TextHeaderCell>Start date</Table.TextHeaderCell>
          <Table.TextHeaderCell>End date</Table.TextHeaderCell>
        </Table.Head>
        <Table.Body>
          {reservations.map((reservation: types.Reservation, i: number) => (
            <Table.Row
              isSelectable
              key={i}
              onSelect={() => alert(reservation.guest.name)}
            >
              <Table.TextCell>{reservation.guest.name}</Table.TextCell>
              <Table.TextCell>{reservation.guest.email}</Table.TextCell>
              <Table.TextCell>{reservation.spaceId}</Table.TextCell>
              <Table.TextCell>{reservation.startDate}</Table.TextCell>
              <Table.TextCell>{reservation.endDate}</Table.TextCell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </Pane>
  );
}

export default ReservationList;
