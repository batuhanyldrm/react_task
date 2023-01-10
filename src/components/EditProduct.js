import React, { useState } from 'react'
import { connect } from 'react-redux';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';
import DialogActions from '@mui/material/DialogActions';
import TextField from '@mui/material/TextField';
import DialogContent from '@mui/material/DialogContent';
import { changeStock } from './api/productApi';
import { updateStock } from './actions/productActions';

function EditProduct(props) {

    const {product, updateStock, open, handleClose, id} = props

    const [productName, setProductName] = useState(product.productName)
    const [description, setDescription] = useState(product.description)
    const [price, setPrice] = useState(product.price)
    const [amount, setAmount] = useState(product.amount)

    const handleChageStock = async () => {
        const data = {
            id: id,
            productName : productName,
            description : description,
            price: price,
            amount: amount,
        }
        await changeStock(data)
        .then(() => {
            updateStock(data)
        }).finally(() =>{
            handleClose(false)
        })
    }

    return(
        <>
         <Dialog
          open={open}
          onClose={handleClose}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
          fullWidth
        >
          <DialogTitle id="alert-dialog-title">
            <div style={{ textAlign : "center" }}>
              Edit Stock
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
            <Button onClick={() =>handleChageStock()}>EDIT</Button>
          </DialogActions>
      </Dialog>
        </>
    )
}

const mapStateToProps = (state) => ({
  });
  
  const mapDispatchToProps = (dispatch) => ({
    updateStock: (data) => {
        dispatch(updateStock(data))
    },
  });

export default connect(mapStateToProps,mapDispatchToProps) (EditProduct)