import React, { useContext } from "react";
import "./FoodDisplay.css";
import { StoreContext } from "../../context/StoreContext";
import FoodItem from "../FoodItem/FoodItem";

const FoodDisplay = ({ category }) => {
  const { food_list } = useContext(StoreContext);
  return (
    <div className="food-display" id="food-display">
      <h2>Size En Yakın Yiyecekler</h2>
      <div className="food-display-list">
        {food_list.map((item, index) => {
          if (category === "All" || category === item.category) {

            return (
              <FoodItem
                key={index}
                id={item.ID}
                name={item.Name}
                price={item.Price}
                image={item.Image}
              />
            );
          }
        })}
      </div>
    </div>
  );
};

export default FoodDisplay;
