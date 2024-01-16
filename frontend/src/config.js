

const API_URL = 'http://localhost:5000';
const WS_URL = 'ws://localhost:5000/monitor';

const wapi = (path='/') => `${API_URL}/${path}`;

export default wapi

export {WS_URL}