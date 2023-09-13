import http from 'k6/http';

export const options ={
    vus:1000,
    iterations:1000000,
};

export default function () {
  const url = 'http://localhost:8080/demo';
  const payload = JSON.stringify({
    text:"sample demo"
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, payload, params);
}