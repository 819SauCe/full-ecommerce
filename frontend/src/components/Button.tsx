import React from "react";

export type ButtonVariant = "primary" | "secondary" | "outline" | "ghost";
export type ButtonSize = "sm" | "md" | "lg";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  variant?: ButtonVariant;
  size?: ButtonSize;
  fullWidth?: boolean;
  loading?: boolean;
}

const Button: React.FC<ButtonProps> = ({ children, variant = "primary", size = "md", fullWidth = false, loading = false, disabled, ...rest }) => {
  const isDisabled = disabled || loading;

  const className = ["btn", variant === "primary" ? "btn-primary" : variant === "secondary" ? "btn-secondary" : variant === "outline" ? "btn-outline-primary" : "btn-link", size === "sm" ? "btn-sm" : size === "lg" ? "btn-lg" : "", fullWidth ? "w-100" : "", rest.className].filter(Boolean).join(" ");

  return (
    <button {...rest} disabled={isDisabled} className={className} type={rest.type || "button"}>
      {loading && <span className="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true" />}
      {children}
    </button>
  );
};

export default Button;
