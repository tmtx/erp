import React from 'react';
import LogInBox from './components/LogInBox'
import { Pane } from 'evergreen-ui';

const App: React.FC = () => {
  return (
    <div className="App">
      <Pane clearfix justifyContent="center" alignItems="center" width="100%" height="100%" display="flex" flexDirection="column" position="relative">
        <LogInBox />
      </Pane>
    </div>
  );
}

export default App;
