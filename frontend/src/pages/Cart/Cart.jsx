import React, { useContext, useEffect, useState } from "react";
import { StoreContext } from "../../context/StoreContext";
import { useNavigate } from "react-router-dom";
import { assets } from "../../assets/assets";
import ErrorPopup from "../../components/ErrorPopup/ErrorPopup";
import SuccessPopup from "../../components/SuccessPopup/SuccessPopup";
import "./Cart.css";

export const deliveryFee = 2;

const Cart = () => {
  const {
    removeFromCart,
    addToCart,
    getTotalCartAmount,
    getTotalQuantity,
    checkout,
    cart_list,
  } = useContext(StoreContext);
  const totalQuantity = getTotalQuantity();
  const navigate = useNavigate();
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const handleCheckout = async () => {
    try {
      await checkout();
      setSuccess("Purchase successful!");
    } catch (err) {
      setError(err.message || "An error occurred during checkout.");
    }
  };

  useEffect(() => {
    
  }, [cart_list]);

  return (
    <div className="cart">
      {error && <ErrorPopup message={error} onClose={() => setError("")} />}
      {success && <SuccessPopup message={success} onClose={() => setSuccess("")} />}
      <div className="cart-items">
        <div className="cart-items-title cart-heading">
          <p>Ürünler</p>
          <p>Ürün İsmi</p>
          <p>Fiyat</p>
          <p>Miktar</p>
          <p>Toplam</p>
          <p> </p>
        </div>
        <br />
        <hr />
        {!cart_list?.length ? (
          <p className="NoItems">Sepette Ürün Yok</p>
        ) : (
          cart_list?.map((item) => (
            <React.Fragment key={item.product_id}>
              <div className="cart-items-title cart-items-item">
                <img src={item.image} alt="food img" />
                <p>{item.name}</p>
                <p>${item.price}</p>
                <div className="quantity-controls">
                  <button onClick={() => removeFromCart(item.product_id)}>-</button>
                  <p>{item.quantity}</p>
                  <button onClick={() => addToCart(item.product_id)}>+</button>
                </div>
                <p>${item.price * item.quantity}</p>
                <p
                  className="Remove"
                  onClick={() => removeFromCart(item.product_id)}
                >
                  <img
                    src={assets.remove_icon_red}
                    alt="remove_icon_red"
                  />
                </p>
              </div>
              <hr key={`hr-${item.product_id}`} />
            </React.Fragment>
          ))
        )}
      </div>
      <div className="cart-bottom">
        <div className="cart-total">
          <h2>Sepet Toplamı</h2>
          <div>
            <div className="cart-total-details">
              <p>Ara Toplam</p>
              <p>${getTotalCartAmount()}</p>
            </div>
            <hr />
            <div className="cart-total-details">
              <p>Servis Ücreti</p>
              <p>${getTotalCartAmount() === 0 ? 0 : deliveryFee}</p>
            </div>
            <hr />
            <div className="cart-total-details">
              <b>Total</b>
              <b>
                ${getTotalCartAmount() === 0
                  ? 0
                  : getTotalCartAmount() + deliveryFee}
              </b>
            </div>
          </div>
          <button
            disabled={getTotalCartAmount() === 0}
            onClick={handleCheckout}
          >
            ÖDEME IŞLEMINE GEÇIN          </button>
        </div>
        <div className="cart-promocode">
          <div>
            {/* Promocode section, if any */}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Cart;
