import axios from "axios";

export const getProducts = async () => {
    const resp = await axios.get("http://localhost:3001/stocks")
    return resp;
}

export const postProduct = async (data) => {
   
        const resp = await axios.post("http://localhost:3001/stocks", {
            productName: data,
            /* description: data,
            price: data,
            amount: data */
        })
        return resp; 
}
