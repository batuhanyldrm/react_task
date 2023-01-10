import React, { useEffect } from 'react'
import { connect } from 'react-redux';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import ProductListItem from './ProductListItem';

function ProductList(props) {

    const {products} = props;


    return(
    <div>
        <TableContainer component={Paper}>
      <Table sx={{ minWidth: 5 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Product Name</TableCell>
            <TableCell align="left">Description</TableCell>
            <TableCell align="right">Price</TableCell>
            <TableCell align="right">Amount</TableCell>
            <TableCell align="right">Delete</TableCell>
            <TableCell align="right">Edit</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {products.products && products.products.map((product, index) => (
            <ProductListItem
            product={product}
            index={index}
            key={product.id + "" + index}
            />
          ))}
        </TableBody>
      </Table>
    </TableContainer>
    </div>
    );
}

const mapStateToProps = (state) => ({
  });
  
  const mapDispatchToProps = (dispatch) => ({
  });

export default connect(mapStateToProps,mapDispatchToProps) (ProductList)