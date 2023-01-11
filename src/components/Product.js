import React, { useState, useEffect } from 'react'
import { connect } from 'react-redux';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import IconButton from '@mui/material/IconButton';
import SearchIcon from '@mui/icons-material/Search';
import AddProduct from './AddProduct';
import ProductList from './ProductList';
import Order from './Order';
import { fetchProducts, fetchSearchProducts } from './actions/productActions';

function Product(props) {

    const {fetchProducts, products, fetchSearchProducts} = props;

    const [open, setOpen] = useState(false);
    const [order, setOrder] = useState(false);
    const [search, setSearch] = useState("");

    const handleSearch = () => {
        fetchSearchProducts(search)
    }

    const checkPressedEnter = (key) => {
        if (key === "Enter") {
            fetchSearchProducts(search)
        }
      };

    useEffect(() => {
        fetchProducts()
      }, [])

    const orderPopUpClose = () => {
        setOrder(false);
    };
    
    const handleClose = () => {
    setOpen(false);
    };

    return(
    <div>
        <Order
            open={order}
            orderPopUpClose={orderPopUpClose}
            products={products}
            fetchProducts={fetchProducts}
        />
        <AddProduct
            open={open}
            handleClose={handleClose}
        />

        <Button variant="contained" color="primary" style={{margin:"5px"}} onClick={() => setOrder(true)}>
            USE PRODUCT
        </Button>
        <Button variant="contained" color="primary" style={{margin:"5px"}} onClick={() => setOpen(true)}>
            ADD PRODUCT
        </Button>
             <TextField 
             style={{marginTop:"5px"}} 
             id="outlined-basic" 
             label="Search" 
             size='small'
             value={search}
             onChange={(e) => setSearch(e.target.value)}
             onKeyPress={(e) => checkPressedEnter(e.key)}
             variant="outlined" 
             InputProps={{
                 endAdornment: (
                     <>
                     <IconButton size="small" onClick={() => handleSearch()}>
                         <SearchIcon/>
                     </IconButton>
                     </>
                 ),
             }}
             />
        <ProductList
            products={products}
        />
    </div>
    );
}

const mapStateToProps = (state) => ({
    products: state.products
  });
  
  const mapDispatchToProps = (dispatch) => ({
    fetchProducts: () => {
      dispatch(fetchProducts());
    },
    fetchSearchProducts: (data) => {
        dispatch(fetchSearchProducts(data));
      },
  });

export default  connect(mapStateToProps,mapDispatchToProps) (Product)