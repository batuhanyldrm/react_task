import React, { useState } from 'react'
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import IconButton from '@mui/material/IconButton';
import SearchIcon from '@mui/icons-material/Search';
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
        <TextField 
        style={{marginTop:"5px"}} 
        id="outlined-basic" 
        label="Search" 
        size='small' 
        variant="outlined" 
        InputProps={{
            endAdornment: (
                <>
                <IconButton size="small">
                    <SearchIcon/>
                </IconButton>
                </>
            ),
        }}
        />
        <ProductList/>
    </div>
    );
}

export default Product