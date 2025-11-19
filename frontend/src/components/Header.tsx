import React, { useEffect, useState } from "react";
import { useNavigate, Link, useLocation } from "react-router-dom";
import { API_ROUTES } from "../config/apiConfig";

type MeResponse = {
  status: string;
  id: any;
  profile_img: any;
  first_name: any;
  last_name: any;
  email: any;
  role: any;
  exp: any;
  iat: any;
};

const Header: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const [me, setMe] = useState<MeResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | undefined>();

  // 🔒 futuramente você pode trocar isso por contexto de carrinho
  const [cartCount] = useState<number>(0);

  const isAuthPage = location.pathname === "/login" || location.pathname === "/register";

  // --- Cache simples pra não gastar request à toa ---
  function loadCachedUser() {
    try {
      const cached = localStorage.getItem("me_data");
      if (!cached) return null;

      const parsed = JSON.parse(cached);
      if (parsed?.exp && parsed.exp * 1000 > Date.now()) {
        return parsed;
      }
      return null;
    } catch {
      return null;
    }
  }

  async function fetchMe() {
    setError(undefined);

    const cached = loadCachedUser();
    if (cached) {
      setMe(cached);
      return;
    }

    try {
      setLoading(true);

      const res = await fetch(API_ROUTES.me, {
        method: "GET",
        credentials: "include",
      });

      if (!res.ok) {
        if (res.status === 401 || res.status === 403) {
          setMe(null);
          return;
        }
        throw new Error("Falha ao carregar usuário.");
      }

      const data: MeResponse = await res.json();
      setMe(data);
      localStorage.setItem("me_data", JSON.stringify(data));
    } catch (err: any) {
      console.error(err);
      setError(err.message || "Não foi possível carregar o usuário.");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    // não precisa buscar /auth/me em telas de login/cadastro
    if (!isAuthPage) {
      fetchMe();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthPage]);

  function handleLogout() {
    localStorage.removeItem("me_data");
    setMe(null);
    navigate("/login");
  }

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-white border-bottom shadow-sm sticky-top">
      <div className="container-fluid px-4">
        {/* Logo */}
        <Link to="/" className="navbar-brand fw-bold fs-3">
          Minha Loja
        </Link>

        {/* Toggle mobile */}
        <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#mainNavbar" aria-controls="mainNavbar" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon" />
        </button>

        <div className="collapse navbar-collapse" id="mainNavbar">
          {/* Links principais (esquerda) */}
          <ul className="navbar-nav me-auto mb-2 mb-lg-0">
            <li className="nav-item">
              <Link to="/categorias" className="nav-link">
                Categorias
              </Link>
            </li>
            <li className="nav-item">
              <Link to="/ofertas" className="nav-link">
                Ofertas
              </Link>
            </li>
            <li className="nav-item">
              <Link to="/novidades" className="nav-link">
                Novidades
              </Link>
            </li>
          </ul>

          {/* Barra de busca menor (central / direita) */}
          {!isAuthPage && (
            <form className="d-none d-md-flex me-3" style={{ maxWidth: 260, width: "100%" }}>
              <div className="input-group input-group-sm">
                <input type="search" className="form-control" placeholder="Buscar..." aria-label="Buscar produtos" />
                <button className="btn btn-outline-secondary" type="submit">
                  Buscar
                </button>
              </div>
            </form>
          )}

          {/* Lado direito: conta + pedidos + carrinho */}
          <ul className="navbar-nav ms-0 ms-md-2 align-items-center gap-2">
            {/* status de loading só como texto pequeno */}
            {loading && !isAuthPage && (
              <li className="nav-item d-none d-lg-flex">
                <span className="small text-muted">Carregando...</span>
              </li>
            )}
            {/* não precisa mostrar erro gigante no header, só se quiser debugar */}
            {error && !isAuthPage && (
              <li className="nav-item d-none d-lg-flex">
                <span className="small text-danger">Erro de sessão</span>
              </li>
            )}

            {/* Conta / Login */}
            {!isAuthPage && (
              <li className="nav-item dropdown">
                {me ? (
                  <>
                    <button className="btn btn-link nav-link dropdown-toggle px-2" id="userDropdown" data-bs-toggle="dropdown" aria-expanded="false" type="button">
                      {/* Mostra só o primeiro nome ou "Minha conta" */}
                      Olá, {me.first_name || "cliente"}
                    </button>
                    <ul className="dropdown-menu dropdown-menu-end" aria-labelledby="userDropdown">
                      <li>
                        <Link to="/minha-conta" className="dropdown-item">
                          Minha conta
                        </Link>
                      </li>
                      <li>
                        <Link to="/pedidos" className="dropdown-item">
                          Meus pedidos
                        </Link>
                      </li>
                      <li>
                        <hr className="dropdown-divider" />
                      </li>
                      <li>
                        <button type="button" className="dropdown-item" onClick={handleLogout}>
                          Sair
                        </button>
                      </li>
                    </ul>
                  </>
                ) : (
                  <button type="button" className="btn btn-outline-primary btn-sm" onClick={() => navigate("/login")}>
                    Entrar
                  </button>
                )}
              </li>
            )}

            {/* Link "Meus pedidos" extra (desktop) */}
            {!isAuthPage && (
              <li className="nav-item d-none d-lg-block">
                <Link to="/pedidos" className="nav-link small">
                  Meus pedidos
                </Link>
              </li>
            )}

            {/* Carrinho */}
            {!isAuthPage && (
              <li className="nav-item">
                <button type="button" className="btn btn-outline-secondary btn-sm position-relative" onClick={() => navigate("/carrinho")}>
                  {/* Se tiver Bootstrap Icons, você pode trocar por <i className="bi bi-cart"></i> */}
                  Carrinho
                  {cartCount > 0 && <span className="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">{cartCount}</span>}
                </button>
              </li>
            )}
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Header;
