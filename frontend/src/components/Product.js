import React from 'react'
import Test from './test.jpeg'
function Product() {
  return (
    <div className='product-mains'>
        <img src={Test} className='product-image' alt='ss'></img>
        <p className='product-name'>Test</p>
        <p className='product-cost'>$12.99</p>
    </div>
  )
}

export default Product