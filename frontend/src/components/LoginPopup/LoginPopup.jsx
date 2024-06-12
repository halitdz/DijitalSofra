import React, { useState, useContext, useEffect } from "react";
import "./LoginPopup.css";
import axios from "axios";
import { assets } from "../../assets/assets";
import { StoreContext } from "../../context/StoreContext";
import ErrorPopup from "../../components/ErrorPopup/ErrorPopup";

const LoginPopup = ({ setShowLogin }) => {
  const [currentState, setCurrentState] = useState("Giriş");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [error, setError] = useState("");
  const [passwordError, setPasswordError] = useState("");

  const { setUser } = useContext(StoreContext);

  useEffect(() => {
    setCurrentState("Giriş");
  }, []);

  const handleSignup = async (e) => {
    e.preventDefault();
    if (!validatePassword(password)) {
      setPasswordError("Şifre en az bir büyük harf, bir küçük harf, bir rakam ve bir özel karakter içermelidir.");
      return;
    }
    try {
      const response = await axios.post("http://localhost:3000/auth/register", {
        Name: name,
        Email: email,
        Password: password,
      });
      setUser(response.data);
      setShowLogin(false);
      window.location.reload();
    } catch (error) {
      setError(error.response?.data?.error || "An error occurred during signup.");
    }
  };

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("http://localhost:3000/auth/login", {
        Email: email,
        Password: password,
      });
      localStorage.setItem("token", response.data.token);
      setUser(response.data.user); // Assuming the response includes a 'user' object
      setShowLogin(false);
      window.location.reload();
    } catch (error) {
      setError(error.response?.data?.error || "An error occurred during login.");
    }
  };

  const validatePassword = (password) => {
    const regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    return regex.test(password);
  };

  return (
    <div className="login-popup">
      {error && <ErrorPopup message={error} onClose={() => setError("")} />}
      <form
        className="login-popup-container"
        onSubmit={currentState === "Kayıt ol" ? handleSignup : handleLogin}
      >
        <div className="login-popup-title">
          <h2>{currentState}</h2>
          <img
            src={assets.cross_icon}
            alt="cross_icon"
            onClick={() => setShowLogin(false)}
          />
        </div>
        <div className="login-popup-inputs">
          {currentState === "Kayıt ol" && (
            <input
              type="text"
              placeholder="Adınız"
              required
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          )}
          <input
            type="email"
            placeholder="Email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <input
            type="password"
            placeholder="Şifre"
            required
            value={password}
            onChange={(e) => {
              setPassword(e.target.value);
              if (passwordError) setPasswordError("");
            }}
          />
          {currentState === "Kayıt ol" && (
            <p className="password-requirements">
              Şifre en az bir büyük harf, bir küçük harf, bir rakam ve bir özel karakter içermelidir.
            </p>
          )}
        </div>

        <button type="submit">
          {currentState === "Kayıt ol" ? "Hesap oluştur" : "Giriş"}
        </button>

        {currentState === "Kayıt ol" && (
          <div className="login-popup-condition">
            <input type="checkbox" required />
            <p>Devam ederek, kullanım koşullarını ve gizlilik politikasını kabul ediyorum.</p>
          </div>
        )}

        {currentState === "Giriş" ? (
          <p>
            Yeni Hesap Oluşturmak İçin
            <span onClick={() => setCurrentState("Kayıt ol")}>Buraya Tıklayın</span>
          </p>
        ) : (
          <p>
            Zaten Hesabınız varsa
            <span onClick={() => setCurrentState("Giriş")}>Buradan Giriş Yapın</span>
          </p>
        )}
      </form>
    </div>
  );
};

export default LoginPopup;
