import React from "react";
import { Routes, Route, useLocation } from "react-router-dom";
import Header from "./components/Header";
import Home from "./pages/Home";
import LoginPage from "./pages/Login";
import RegisterPage from "./pages/Register";
import NotFound from "./pages/NotFound";

const App: React.FC = () => {
  const location = useLocation();

  const hideHeaderOn = ["/login", "/register"];
  const shouldShowHeader = !hideHeaderOn.includes(location.pathname);

  return (
    <>
      {shouldShowHeader && <Header />}

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
};

export default App;
