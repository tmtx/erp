import React from 'react';
import { Pane, Button, Heading, TextInput } from 'evergreen-ui';

const App: React.FC = () => {
  return (
    <Pane marginTop="15%" display="flex" alignItems="center" flexDirection="column">
      <Heading size={800}>Authenticate</Heading>
      <Pane
        elevation={3}
        backgroundColor="white"
        width={300}
        height={150}
        margin={24}
        display="flex"
        justifyContent="center"
        alignItems="center"
        flexDirection="column"
      >
        <TextInput
          marginTop={15}
          marginBottom={10}
          name="email"
          placeholder="Email"
        />
        <TextInput
          marginBottom={15}
          name="password"
          placeholder="Password"
        />
        <Button>Log in</Button>
      </Pane>
    </Pane>
  );
}

export default App;
