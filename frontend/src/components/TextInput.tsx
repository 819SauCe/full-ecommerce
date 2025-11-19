// src/components/ui/TextInput.tsx
import React from "react";

interface TextInputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
}

const TextInput: React.FC<TextInputProps> = ({ label, error, id, ...rest }) => {
  const inputId = id || rest.name;

  return (
    <div className="mb-3">
      {label && (
        <label htmlFor={inputId} className="form-label">
          {label}
        </label>
      )}
      <input id={inputId} className={`form-control ${error ? "is-invalid" : ""}`} {...rest} />
      {error && <div className="invalid-feedback">{error}</div>}
    </div>
  );
};

export default TextInput;
