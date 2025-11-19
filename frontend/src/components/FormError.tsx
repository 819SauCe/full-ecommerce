// src/components/ui/FormError.tsx
import React from "react";

interface FormErrorProps {
  message?: string;
}

const FormError: React.FC<FormErrorProps> = ({ message }) => {
  if (!message) return null;
  return (
    <div className="alert alert-danger py-2" role="alert">
      {message}
    </div>
  );
};

export default FormError;
