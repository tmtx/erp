import React from "react";
import { Menu, Popover, Position } from "evergreen-ui";

import API from "./../Api";

type Props = {
  children: React.ReactNode,
  onProfileEditOpen: () => void,
};
const HeaderMenu = (props: Props) => {
  const logOut = () => {
    API.post("/users/logout")
      .then( response => {
        if (response.data && response.data.status === "ok") {
          window.location.reload();
        }
      });
  };

  return (
    <Popover
      position={Position.BOTTOM_LEFT}
      content={
        <Menu>
          <Menu.Group>
            <Menu.Item icon="person" onSelect={props.onProfileEditOpen}>
              Edit info
            </Menu.Item>
          </Menu.Group>
          <Menu.Divider />
          <Menu.Group>
            <Menu.Item icon="power" intent="danger" onSelect={ () => logOut() }>
              Logout
            </Menu.Item>
          </Menu.Group>
        </Menu>
      }
    >
      { props.children }
    </Popover>
  );
};

export default HeaderMenu;
