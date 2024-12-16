import axios    from "axios";

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;
const service = axios.create({
    baseURL:  '/api',
    timeout: 10000,
})



export default service;