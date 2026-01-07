export interface ApiResponse<T> {
  payload: {
    data: T;
  };
  messages: string[];
  errorCode: number;
}

