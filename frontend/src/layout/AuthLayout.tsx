// src/components/layout/AuthLayout.tsx
import React from "react";

interface AuthLayoutProps {
  title: string;
  subtitle?: string;
  children: React.ReactNode;
}

const AuthLayout: React.FC<AuthLayoutProps> = ({ title, subtitle, children }) => {
  return (
    <div className="min-vh-100 d-flex align-items-center justify-content-center bg-light">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-12 col-sm-10 col-md-8 col-lg-5">
            <div className="card shadow-sm border-0">
              <div className="card-body p-4 p-md-5">
                <h1 className="h3 mb-2 text-center">{title}</h1>
                {subtitle && <p className="text-muted text-center mb-4">{subtitle}</p>}
                {children}
              </div>
            </div>
            <p className="text-center text-muted mt-3" style={{ fontSize: 12 }}>
              © {new Date().getFullYear()} - Seu E-commerce
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default AuthLayout;
