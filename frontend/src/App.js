import './App.css';
import Footer from './components/Footer';
import Home from './components/Home';
import Navbar from './components/Navbar';
import Products from './components/Products';
import Categories from './components/Categories';
import Deals from './components/Deals';
import About from './components/About';
import Contact from './components/Contact';
import {BrowserRouter , Routes, Route} from "react-router-dom";

function App() {
  return (
    <div className="App">
     <BrowserRouter>
     <Navbar />
     <Routes>
      <Route path="/" element={<Home />}></Route>
      <Route path="/products" element={<Products />}></Route>
      <Route path="/categories" element={<Categories />}></Route>
      <Route path="/deals" element={<Deals />}></Route>
      <Route path="/about" element={<About />}></Route>
      <Route path="/Contact" element={<Contact />}></Route>
     </Routes>
     <Footer />
     </BrowserRouter>
    </div>
  );
}

export default App;
