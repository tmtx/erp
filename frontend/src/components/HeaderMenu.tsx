import React from "react";
import { Menu, Popover, Position } from "evergreen-ui";

type Props = {
  children: React.ReactNode,
  onProfileEditOpen: () => void,
};
const HeaderMenu = (props: Props) => {
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
            <Menu.Item icon="power" intent="danger">
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
