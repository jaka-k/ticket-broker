type Status = 'outbound' | 'pending' | 'canceled' | 'confirmed' | 'rejected';

type Order = {
  id: string;
  countryCode: string;
  orderAmount: number;
  status: Status;
};
