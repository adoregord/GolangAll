import http from 'k6/http';
import { sleep } from 'k6';
import { check } from 'k6';

export const options = {
    vus: 100, // Number of virtual users
    duration: '10s', // Duration of the test
};

const payload = JSON.stringify({
    "userid": 1,
    "eventid": 7,
    "ticket": [
        {"ID": 1, "Type": "VIP", "Quantity": 1},
        {"ID": 2, "Type": "CAT 1", "Quantity": 1}
    ]
});

const params = {
    headers: {
        'Content-Type': 'application/json',
    },
};

export default function () {
    const res = http.post('http://localhost:8080/buyTicket', payload, params);
    check(res, {
        'status is 200': (r) => r.status === 200,
    });
    sleep(1);
}
