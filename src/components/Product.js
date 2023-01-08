import React, { useState } from 'react'
import Button from '@mui/material/Button';
import AddProduct from './AddProduct';
import ProductList from './ProductList';

function Product(props) {

    const [open, setOpen] = useState(false);
    
    const handleClose = () => {
    setOpen(false);
    };

    return(
    <div>
        <AddProduct
            open={open}
            handleClose={handleClose}
        />

        <Button variant="contained" color="primary" style={{margin:"5px"}} onClick={() => setOpen(true)}>
            ADD PRODUCT
        </Button>
        <ProductList/>
    </div>
    );
}

export default Product