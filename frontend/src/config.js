
// const API_URL = 'http://localhost:5000';
// const WS_URL = 'ws://localhost:5000/monitor';

const API_URL = 'https://bread.m41den.com';
const WS_URL = 'wss://bread.m41den.com/monitor';

const wapi = (path='/') => `${API_URL}/${path}`;

export default wapi

export {WS_URL}