import React from 'react';

import * as types from './types';

export const UserContext = React.createContext<types.User|null>({
  email: null,
  id: null,
});
