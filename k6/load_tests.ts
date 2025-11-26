// @ts-ignore
import http from 'k6/http';
// @ts-ignore
import { sleep } from 'k6';

export const options = {
    iterations: 1,
    // stages: [
    //     { duration: '30s', target: 20 },
    //     { duration: '1m30s', target: 10 },
    //     { duration: '20s', target: 0 },
    // ],
}

export default function () {
    // Make a GET request to the target URL
    http.post('http://localhost:8080/solve');

    console.log("Running...");

    // Sleep for 1 second to simulate real-world usage
    sleep(1);
}