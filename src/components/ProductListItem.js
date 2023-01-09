import React, { useState } from 'react'
import { connect } from 'react-redux';
import TableCell from '@mui/material/TableCell';
import TableRow from '@mui/material/TableRow';
import DeleteIcon from '@mui/icons-material/Delete';
import IconButton from '@mui/material/IconButton';
import EditIcon from '@mui/icons-material/Edit';
import { deleteProduct } from './actions/productActions';
import EditProduct from './EditProduct';

function ProductListItem(props) {

    const {product, deleteProduct} = props

    const [id,setId] = useState("")
    const [open, setOpen] = useState(false);
    const handleClose = () => setOpen(false);

    const handleEdit = (ID) => {
      setId(ID)
      setOpen(true);
    }

    return(
        <>
        <EditProduct
          open={open}
          handleClose={handleClose}
          product={product}
          id={id}
        />
        <TableRow
          key={product.id}
          sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
        >
            
          <TableCell component="th" scope="row" align="left">
            {product.productName}
          </TableCell>
          <TableCell align="left">{product.description}</TableCell>
          <TableCell align="right">{product.price}</TableCell>
          <TableCell align="right">{product.amount}</TableCell>
          <TableCell align="right">
            <IconButton
            onClick={()=>deleteProduct(product.id)}
            >
                <DeleteIcon/>
            </IconButton>
          </TableCell>
          <TableCell align="right">
          <IconButton
          onClick={()=>handleEdit(product.id)}
          >
                <EditIcon/>
            </IconButton>
          </TableCell>
        </TableRow>
        </>
    )
}

const mapStateToProps = (state) => ({
  });
  
  const mapDispatchToProps = (dispatch) => ({
    deleteProduct: (id) => {
        dispatch(deleteProduct(id));
      },
  });

export default connect(mapStateToProps,mapDispatchToProps) (ProductListItem)