// src/components/ui/Checkbox.tsx
import React from "react";

interface CheckboxProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label: string;
}

const Checkbox: React.FC<CheckboxProps> = ({ label, id, ...rest }) => {
  const inputId = id || rest.name;

  return (
    <div className="form-check">
      <input id={inputId} type="checkbox" className="form-check-input" {...rest} />
      <label className="form-check-label" htmlFor={inputId}>
        {label}
      </label>
    </div>
  );
};

export default Checkbox;
