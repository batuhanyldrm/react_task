import { ADD_PRODUCT, FETCH_PRODUCTS } from "../actions/types";

const Reducer = (state = {}, action) => {
    switch (action.type) {
        case ADD_PRODUCT:
            return {...state, allproducts:  [...state.allproducts, action.payload]}
        case FETCH_PRODUCTS:
            return {...state, products: action.payload}
        default:
            return state
    }
};

export default Reducer;
