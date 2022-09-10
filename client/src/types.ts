export type OrderData = {
  dish: string;
  server: string;
  table: string;
  price: string;
}

export type OrderId = {
  _id: number;
}

export type ChangeOrderOrWaiter = {
  change: boolean;
  id: number;
}