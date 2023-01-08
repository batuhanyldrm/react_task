import React, { useState } from 'react';
import { connect } from 'react-redux';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';
import DialogActions from '@mui/material/DialogActions';
import TextField from '@mui/material/TextField';
import { addProduct } from './actions/productActions';
import { postProduct } from './api/productApi';

function AddProduct(props) {

    const {open, handleClose, addProduct} = props

    const [product, setProduct] = useState("")

    const handleCreateProduct = async (data) => {
      await postProduct(data
        ).then((res) => {
        addProduct(res.data)
      })
    }

    return(
    <div>
         <Dialog
          open={open}
          onClose={handleClose}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
          fullWidth
        >
          <DialogTitle id="alert-dialog-title">
            <div style={{ textAlign : "center" }}>
              Add Stock
            </div>
          </DialogTitle>
          <TextField
            id="product"
            value={product}  
            onChange={(e) => setProduct(e.target.value)}
            type="text"
            margin="dense"
            label="Product Name" 
            variant="outlined"
            size='small'
          />
         {/*  <TextField
            value={product}
            onChange={(e) => setProduct(e.target.value)}
            autoFocus
            margin="dense"
            id="data"
            label="Todo"
            type="text"
            fullWidth
            variant="standard"
          /> */}
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button onClick={() =>handleCreateProduct(product)}>ADD</Button>
          </DialogActions>
      </Dialog>
    </div>
    );
}

const mapStateToProps = (state) => ({
});

const mapDispatchToProps = (dispatch) => ({
  addProduct: (data) => {
    dispatch(addProduct(data));},
  });

export default connect(mapStateToProps, mapDispatchToProps) (AddProduct)