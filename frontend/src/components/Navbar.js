import React from 'react'
import { useNavigate } from 'react-router-dom'

function Navbar() {

    const navigate = useNavigate();

    const handleClick = (path) => {
        navigate(path);
    }

  return (
    <nav className="navbar">
    <div className="logo">WebMarket</div>
    <ul className="nav-links">
        <li onClick={() => {handleClick('/')}}>Home</li>
        <li onClick={() => {handleClick('/products')}}>Products</li>
        <li className="dropdown" onClick={() => {handleClick('/categories')}}>
           Categories
            <div className="dropdown-content">
           
            </div>
        </li>
        <li onClick={() => {handleClick('/deals')}}>Deals</li>
        <li onClick={() => {handleClick('/about')}}>About Us</li>
        <li onClick={() => {handleClick('/contact')}}>Contact</li>
    </ul>
    <div className="burger">
        <div className="line1"></div>
        <div className="line2"></div>
        <div className="line3"></div>
    </div>
</nav>
  )
}

export default Navbar