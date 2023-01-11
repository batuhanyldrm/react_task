import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';
import DialogActions from '@mui/material/DialogActions';
import TextField from '@mui/material/TextField';
import DialogContent from '@mui/material/DialogContent';
import { addProduct } from './actions/productActions';
import { postProduct } from './api/productApi';

function AddProduct(props) {

    const {open, handleClose, addProduct} = props

    const [productName, setProductName] = useState("")
    const [description, setDescription] = useState("")
    const [price, setPrice] = useState(0)
    const [amount, setAmount] = useState(0)

    const handleCreateProduct = async () => {
      const data = {
        productName : productName,
        description : description,
        price: price,
        amount: amount,
      }
      await postProduct(data
        ).then((res) => {
        addProduct(res.data)
      }).finally(() => {
        handleClose(false)
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
          <DialogContent>
            <div style={{display:"grid"}}>
          <TextField
            id="product"
            value={productName}
            onChange={(e) => setProductName(e.target.value)}
            type="text"
            margin="normal"
            label="Product Name"
            variant="outlined"
            size='small'
          />
          <TextField
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            autoFocus
            margin="normal"
            id="description"
            label="Description"
            type="text"
            variant="outlined"
            size='small'
          />
          <TextField
            value={amount}
            onChange={(e) => setAmount(parseInt(e.target.value))}
            autoFocus
            margin="normal"
            id="amount"
            label="Amount"
            type="number"
            variant="outlined"
            size='small'
          />
          <TextField
            value={price}
            onChange={(e) => setPrice(parseInt(e.target.value))}
            autoFocus
            margin="normal"
            id="price"
            label="Price"
            type="number"
            variant="outlined"
            size='small'
          />
          </div>
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button onClick={() =>handleCreateProduct()}>ADD</Button>
          </DialogActions>
      </Dialog>
    </div>
    );
}

const mapStateToProps = (state) => ({
});

const mapDispatchToProps = (dispatch) => ({
  addProduct: (data) => {
    dispatch(addProduct(data));
  },
});

export default connect(mapStateToProps, mapDispatchToProps) (AddProduct)