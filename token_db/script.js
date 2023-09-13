import http from 'k6/http';

export const options ={
    vus:100,
    iterations:1000000,
};

export default function () {
  const url = 'http://localhost:4000/store';
  const payload = JSON.stringify({
    // Your JSON payload here
    });

  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2OTQ1ODEwNjQsIm5hbWUiOiJzdWJhc3JpIiwicGFzc3dvcmQiOiJzdWJhMTIzNCIsInN1YiI6IjEyMzQ1Njc4OTAifQ.i7cVK5s6Q-x_RsrrQfqh81MGtnUSiBYV2rwh7CpBn4U'
    },
  };

  http.post(url,payload, params);
}