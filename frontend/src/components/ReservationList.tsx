import React from "react";
import { Table, Pane, Button } from "evergreen-ui";
import { Link } from "react-router-dom";

const ReservationList: React.FC = () => {
  let reservations = [
    {
      id: 2,
      guest: {
        name: "Kārlis Feldmanis",
        email: "kf@karlis.dev",
      },
      startDate: "2019-11-11",
      endDate: "2019-12-12",
      livingSpaceId: 4,
    },
    {
      id: 4,
      guest: {
        name: "Jānis Bērziņš",
        email: "janis@karlis.dev",
      },
      startDate: "2020-01-11",
      endDate: "2020-02-12",
      livingSpaceId: 5,
    },
    {
      id: 5,
      guest: {
        name: "Andris T.",
        email: "andris@karlis.dev",
      },
      startDate: "2020-03-11",
      endDate: "2020-04-22",
      livingSpaceId: 4,
    },
    {
      id: 6,
      guest: {
        name: "Edgars",
        email: "edgars@karlis.dev",
      },
      startDate: "2020-06-11",
      endDate: "2020-07-12",
      livingSpaceId: 4,
    },
  ];

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
          <Table.SearchHeaderCell placeholder="Name" />
          <Table.TextHeaderCell>Email</Table.TextHeaderCell>
          <Table.TextHeaderCell>Room number</Table.TextHeaderCell>
          <Table.TextHeaderCell>Start date</Table.TextHeaderCell>
          <Table.TextHeaderCell>End date</Table.TextHeaderCell>
        </Table.Head>
        <Table.Body>
          {reservations.map(reservation => (
            <Table.Row
              isSelectable
              key={reservation.id}
              onSelect={() => alert(reservation.guest.name)}
            >
              <Table.TextCell>{reservation.guest.name}</Table.TextCell>
              <Table.TextCell>{reservation.guest.email}</Table.TextCell>
              <Table.TextCell>{reservation.livingSpaceId}</Table.TextCell>
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
