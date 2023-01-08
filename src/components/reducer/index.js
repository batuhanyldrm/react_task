import {combineReducers} from "redux";
import productReducer from "./productReducer";

const reducers = combineReducers({
    product: productReducer,
});

export default reducers;