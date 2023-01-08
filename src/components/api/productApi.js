import axios from "axios";

export const getProducts = async () => {
    const resp = await axios.get("http://localhost:3001/stocks")
    return resp;
}

export const postProduct = async ({productName,description,price,amount}) => {
   
        const resp = await axios.post("http://localhost:3001/stocks", {
            productName: productName,
            description: description,
            price: price,
            amount: amount,
        })
        return resp; 
}

export const removeProduct = async (id) => {
    const resp = await axios.delete(`http://localhost:3001/stocks/${id}`)
    return resp;
}
