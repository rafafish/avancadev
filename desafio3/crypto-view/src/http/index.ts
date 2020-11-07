import axios from 'axios';

const apiKey = '70381000a6d452de7ddbcd6cafe8ceaa89e3ed7dedac1a550166622b1f489bf4'

export const cryptoHttp = axios.create({
    baseURL: 'https://min-api.cryptocompare.com/data',
    headers:{
        authorization: `Apikey ${apiKey}`
    }
});