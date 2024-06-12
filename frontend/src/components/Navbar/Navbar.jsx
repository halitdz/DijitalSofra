import React, { useContext } from "react";
import { StoreContext } from "../../context/StoreContext";
import { useNavigate, useLocation } from "react-router-dom";
import "./Navbar.css";
import { assets } from "../../assets/assets";

const Navbar = ({ setShowLogin }) => {
  const { user, logout } = useContext(StoreContext);
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogout = () => {
    logout();
    navigate("/");
    window.location.reload(); // Sayfayı yenile
  };

  return (
    <div className="navbar">
      <div className="navbar-logo" onClick={() => navigate("/")}>
        <img src={assets.logo} alt="logo" />
      </div>
      <div className="navbar-links">
        <button 
          className={location.pathname === "/" ? "active" : ""} 
          onClick={() => navigate("/")}
        >
         Anasayfa
        </button>
        <button 
          className={location.pathname === "/cart" ? "active" : ""} 
          onClick={() => navigate("/cart")}
        >
          Sepet
        </button>
        {!user ? (
          <button onClick={() => setShowLogin(true)}>Giriş</button>
        ) : (
          <>
            <span className="balance">Balance: ${user.balance}</span>
            <button className="logout-button" onClick={handleLogout}>Çıkış</button>
          </>
        )}
      </div>
    </div>
  );
};

export default Navbar;
