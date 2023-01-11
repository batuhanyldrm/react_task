import { getProducts, removeProduct, updateProductAmount, getSearch } from "../api/productApi";
import { ADD_PRODUCT, FETCH_PRODUCTS, DELETE_PRODUCT, UPDATE_STOCK, UPDATE_PRODUCT_STOCK, FETCH_SEARCH_PRODUCT } from "./types";

export const fetchProducts = () => async (
    dispatch
) => {
    const resp = await getProducts()
        dispatch({
            type: FETCH_PRODUCTS,
            payload: resp.data
        })
    
}

export const fetchSearchProducts = (data) => async (
    dispatch
) => {
    const resp = await getSearch(data)
        dispatch({
            type: FETCH_SEARCH_PRODUCT,
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

export const deleteProduct = (id) => async (
    dispatch
) => {
    const resp = await removeProduct(id)
        dispatch({
            type: DELETE_PRODUCT,
            payload: id
        })
}

export const updateStock = (data) => async (
    dispatch
) => {
    dispatch({
        type: UPDATE_STOCK,
        payload: data
    })
}

export const updateProductStock = (id, amount) => async (
    dispatch
) => {
        dispatch({
            type: UPDATE_PRODUCT_STOCK,
            payload: {id,amount}
        })
}
