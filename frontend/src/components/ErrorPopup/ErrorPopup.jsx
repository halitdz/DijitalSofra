import React from 'react';
import './ErrorPopup.css';

const ErrorPopup = ({ message, onClose }) => {
  return (
    <div className="error-popup">
      <div className="error-popup-content">
        <p>{message}</p>
        <button onClick={onClose}>Close</button>
      </div>
    </div>
  );
};

export default ErrorPopup;
