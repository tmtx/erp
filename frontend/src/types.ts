export interface User {
  email: string|null;
  id: string|null;
}

export interface ValidationMessages {
  [index: string]: string;
};

export interface Guest {
  name: string;
  email: string;
}

export interface Reservation {
  spaceId: number;
  guest: Guest;
  startDate: string;
  endDate: string;
}

export interface Space {
  id: number;
}
