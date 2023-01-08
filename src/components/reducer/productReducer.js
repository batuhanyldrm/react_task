import { ADD_PRODUCT, FETCH_PRODUCTS, DELETE_PRODUCT } from "../actions/types";

const Reducer = (state = {}, action) => {
    switch (action.type) {
        case ADD_PRODUCT:
            return {...state, products:  [...state.products, action.payload]};
        case FETCH_PRODUCTS:
            return {...state, products: action.payload}
        case DELETE_PRODUCT:
            return{...state, products: [...state.products].filter((item) => item.id !== action.payload)}
        default:
            return state
    }
};

export default Reducer;
