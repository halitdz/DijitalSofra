import { createContext, useEffect, useState } from "react";
import axios from "axios";

export const StoreContext = createContext(null);

const StoreContextProvider = (props) => {
  const [cartItems, setCartItems] = useState(() => {
    const savedCartItems = localStorage.getItem("cartItems");
    return savedCartItems ? JSON.parse(savedCartItems) : {};
  });
  const [food_list, setFoodList] = useState([]);
  const [cart_list, setCartList] = useState([]);
  const [user, setUser] = useState(null);

  useEffect(() => {
    fetchProducts();
    fetchCartItems();
    fetchUser();
  }, []);

  const fetchProducts = async () => {
    try {
      const response = await axios.get("http://localhost:3000/product/products");
      setFoodList(response.data);
    } catch (error) {
      console.error("Error fetching products:", error);
    }
  };

  const fetchCartItems = async () => {
    try {
      const savedCartItems = JSON.parse(localStorage.getItem("cartItems"));
      if (savedCartItems) {
        const response = await axios.post("http://localhost:3000/cart", { cartItems: savedCartItems });
        
        setCartList(response.data);
      } else {
        setCartList([]); // Ensure cart_list is set to an empty array if there are no saved cart items
      }
    } catch (error) {
      console.error("Error fetching cart items:", error);
    }
  };

  const fetchUser = async () => {
    try {
      const token = localStorage.getItem("token");
      if (!token) throw new Error("You need to log in.");
      const response = await axios.get("http://localhost:3000/user", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setUser(response.data);
    } catch (error) {
      console.error("Error fetching user data:", error);
    }
  };

  const addToCart = (itemId) => {
    setCartItems((prev) => {
      const newCartItems = { ...prev };
      if (newCartItems[itemId]) {
        newCartItems[itemId] += 1;
      } else {
        newCartItems[itemId] = 1;
      }
      localStorage.setItem("cartItems", JSON.stringify(newCartItems));
      fetchCartItems(); // Update cart_list after adding to cart
      return newCartItems;
    });
  };

  const removeFromCart = (itemId) => {
    setCartItems((prev) => {
      const newCartItems = { ...prev };
      if (newCartItems[itemId] === 1) {
        delete newCartItems[itemId];
      } else if (newCartItems[itemId]) {
        newCartItems[itemId] -= 1;
      }
      localStorage.setItem("cartItems", JSON.stringify(newCartItems));
      fetchCartItems(); // Update cart_list after removing from cart
      return newCartItems;
    });
  };

  const getTotalQuantity = () => {
    return Object.values(cartItems)?.reduce((total, quantity) => total + quantity, 0);
  };

  const getTotalCartAmount = () => {
    if (!cart_list) {
      return 0;
    }
    return cart_list?.reduce((total, item) => total + item.price * item.quantity, 0);
  };

  const checkout = async () => {
    try {
      const token = localStorage.getItem("token");
      if (!token) throw new Error("You need to log in.");
      
      const userResponse = await axios.get("http://localhost:3000/user/", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const user = userResponse.data;
      if (!user) throw new Error("User not found. Please register.");

      const totalAmount = getTotalCartAmount();

      if (user.balance < totalAmount) {
        throw new Error("Insufficient balance for checkout.");
      }

      await axios.post(
        "http://localhost:3000/cart/checkout",
        { cartItems },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      // Clear cart after successful checkout
      setCartItems({});
      setCartList([]);
      localStorage.removeItem("cartItems"); // Clear localStorage
      fetchUser(); // Fetch updated user data after checkout
    } catch (error) {
      console.error("Error during checkout:", error);
      throw error; // Rethrow the error to handle it in the Cart component
    }
  };

  const logout = () => {
    localStorage.removeItem("token");
    setUser(null);
  };

  const contextValue = {
    food_list,
    cartItems,
    setCartItems,
    addToCart,
    removeFromCart,
    getTotalCartAmount,
    getTotalQuantity,
    setUser,
    user,
    checkout,
    cart_list, // Add cart_list to context value
    logout,
  };

  return (
    <StoreContext.Provider value={contextValue}>
      {props.children}
    </StoreContext.Provider>
  );
};

export default StoreContextProvider;
