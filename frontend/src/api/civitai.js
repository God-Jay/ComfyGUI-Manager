import axios from "axios";

// import { HttpsProxyAgent } from 'https-proxy-agent';

// const agent = new HttpsProxyAgent('http://127.0.0.1:1083');

export const getCivitaiInfo = async (sha256) => {
    try {
        axios.defaults.timeout = 10000;
        const {data} = await axios.get('https://civitai.com/api/v1/model-versions/by-hash/' + sha256, {
            // httpsAgent: agent
        })
        return data;
    } catch (error) {
        console.error(error);
        throw error;
    }
}