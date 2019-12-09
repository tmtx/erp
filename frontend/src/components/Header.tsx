import React, { useState } from "react";
import { Pane, Avatar, SideSheet, IconButton } from "evergreen-ui";
import { Link } from "react-router-dom";

import HeaderMenu from "./HeaderMenu";
import ProfileEditForm from "./ProfileEditForm";

interface User {
  email: string|null;
}

type Props = { user: User|null } & typeof defaultProps;
const defaultProps = {
  user: null,
}
const Header = (props: Props) => {
  const [showProfileEdit, setShowProfileEdit] = useState(false);

  const getName = (user: User|null): string|null => {
    if (user) {
      return user.email;
    }

    return "";
  };

  return (
    <Pane
      height={50}
      width="100%"
      border="default"
      position="relative"
      display="flex"
      justifyContent="flex-end"
      alignItems="center"
    >
      <div
        style={{
          marginLeft: "auto",
          marginRight: "auto",
          height: "100%",
          width: "800px",
          display: "flex"
        }}
      >
        <IconButton
          is={Link}
          to="/"
          icon="home"
          appearance="minimal"
          height="100%"
          width={50}
        />
      </div>
      <div style={{ height: "100%", display: "flex", width: "200px" }}>
        <HeaderMenu
          onProfileEditOpen={ () => { setShowProfileEdit(true) } }
        >
          <Avatar
            name={getName(props.user)}
            size={40}
            marginTop="auto"
            marginBottom="auto"
            marginRight={50}
            cursor="pointer"
            onClick={ (e: Event) => {} }
          />
        </HeaderMenu>
      </div>
      <SideSheet
        isShown={showProfileEdit}
        onCloseComplete={ () => setShowProfileEdit(false) }
      >
        <ProfileEditForm />
      </SideSheet>
    </Pane>
  );
};

Header.defaultProps = defaultProps;
export default Header;
