import React, { useState } from 'react';
import { connect } from 'react-redux';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';
import DialogActions from '@mui/material/DialogActions';
import TextField from '@mui/material/TextField';
import DialogContent from '@mui/material/DialogContent';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import { updateProductAmount } from './api/productApi';
import { updateProductStock } from './actions/productActions';

function Order(props) {

    const {products, open, updateProductStock, orderPopUpClose} = props

    const [selectedProduct, setSelectedProduct] = useState({})
    const [amount, setAmount] = useState(0)

    const handleSave = async () => {
        if (amount > selectedProduct.amount) {
            console.log("hata")
        }
        else{
           await updateProductAmount(selectedProduct.id,amount)
            .then(() => {
                updateProductStock(selectedProduct.id,amount)
                console.log(updateProductStock(selectedProduct.id,amount),selectedProduct.id,amount)
            }).finally(() => {
                orderPopUpClose(false)
            })
        }
    }


    return(
    <div>
        <Dialog
          open={open}
          onClose={orderPopUpClose}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
          fullWidth
        >
          <DialogTitle id="alert-dialog-title">
            <div style={{ textAlign : "center" }}>
              Use Product
            </div>
          </DialogTitle>
          <DialogContent>
            <div style={{display:"grid"}}>
                    <FormControl variant="standard" sx={{ m: 1, minWidth: 120 }}>
                    <InputLabel id="demo-simple-select-standard-label">Product Name</InputLabel>
                    <Select
                      labelId="demo-simple-select-standard-label"
                      id="demo-simple-select-standard"
                      value={selectedProduct.productName || ""}
                      label="Product Name"
                    >
                    {products.products && products.products.map((product) => {
                       return <MenuItem value={product.productName} onClick={() => setSelectedProduct(product)}>{product.productName}</MenuItem>
                    })}
                    </Select>
                  </FormControl>
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
          </div>
          </DialogContent>
          <DialogActions>
            <Button onClick={orderPopUpClose}>Cancel</Button>
            <Button onClick={() => handleSave()}>Save</Button>
          </DialogActions>
      </Dialog>
    </div>
    );
}

const mapStateToProps = (state) => ({
});

const mapDispatchToProps = (dispatch) => ({
    updateProductStock: (data) => {
        dispatch(updateProductStock(data))
    },
});

export default connect(mapStateToProps, mapDispatchToProps) (Order)