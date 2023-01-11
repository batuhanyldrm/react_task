import { ADD_PRODUCT, FETCH_PRODUCTS, DELETE_PRODUCT, UPDATE_STOCK, UPDATE_PRODUCT_STOCK, FETCH_SEARCH_PRODUCT } from "../actions/types";

const Reducer = (state = {}, action) => {
    switch (action.type) {
        case ADD_PRODUCT:
            return {...state, products:  [...state.products, action.payload]};
        case FETCH_PRODUCTS:
            return {...state, products: action.payload}
        case FETCH_SEARCH_PRODUCT:
            return {...state, products: action.payload}
        case DELETE_PRODUCT:
            return{...state, products: [...state.products].filter((item) => item.id !== action.payload)}
        case UPDATE_STOCK:
            var temp={...state};
            temp.products.map((item, index) => {
                if(item.id == action.payload.id) {
                    temp.products[index].productName = action.payload.productName
                    temp.products[index].description = action.payload.description
                    temp.products[index].price = action.payload.price
                    temp.products[index].amount = action.payload.amount
                }
            })
            return{...state, products: temp.products}
            case UPDATE_PRODUCT_STOCK:
                var temp={...state};
                temp.products.map((item, index) => {
                    if(item.id == action.payload.id) {
                        temp.products[index].amount = item.amount - action.payload.amount
                    }
                })
                return{...state, products: temp.products}
        default:
            return state
    }
};

export default Reducer;
