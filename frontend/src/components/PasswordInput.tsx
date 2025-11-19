// src/components/ui/PasswordInput.tsx
import React, { useState } from "react";

interface PasswordInputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
}

const PasswordInput: React.FC<PasswordInputProps> = ({ label, error, id, ...rest }) => {
  const [visible, setVisible] = useState(false);
  const inputId = id || rest.name;

  return (
    <div className="mb-3">
      {label && (
        <label htmlFor={inputId} className="form-label">
          {label}
        </label>
      )}
      <div className="input-group">
        <input id={inputId} type={visible ? "text" : "password"} className={`form-control ${error ? "is-invalid" : ""}`} {...rest} />
        <button type="button" className="btn btn-outline-secondary" onClick={() => setVisible((v) => !v)} tabIndex={-1}>
          {visible ? "Ocultar" : "Mostrar"}
        </button>
        {error && <div className="invalid-feedback d-block">{error}</div>}
      </div>
    </div>
  );
};

export default PasswordInput;
