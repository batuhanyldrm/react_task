import { getProducts } from "../api/productApi";
import { ADD_PRODUCT, FETCH_PRODUCTS } from "./types";

export const fetchProducts = () => async (
    dispatch
) => {
    const resp = await getProducts()
        dispatch({
            type: FETCH_PRODUCTS,
            payload: resp.data
        })
    
}

export const addProduct = (product) => async (
    dispatch
) => {
    dispatch({
        type: ADD_PRODUCT,
        payload: product
    })
}