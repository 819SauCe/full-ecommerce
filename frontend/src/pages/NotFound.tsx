import { useEffect, useState, useMemo } from "react";
import { useNavigate } from "react-router-dom";
import "../styles/NotFound.scss";

const REDIRECT_SECONDS = 15;

export default function NotFound() {
  const navigate = useNavigate();
  const [counter, setCounter] = useState(REDIRECT_SECONDS);

  useEffect(() => {
    const interval = setInterval(() => {
      setCounter((prev) => {
        if (prev <= 1) {
          return 0;
        }
        return prev - 1;
      });
    }, 1000);
    const timer = setTimeout(() => {
      navigate("/");
    }, REDIRECT_SECONDS * 1000);

    return () => {
      clearInterval(interval);
      clearTimeout(timer);
    };
  }, [navigate]);

  const progress = useMemo(() => (counter / REDIRECT_SECONDS) * 100, [counter]);

  return (
    <div className="d-flex align-items-center justify-content-center min-vh-100 bg-light">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-lg-6 col-md-8">
            <div className="card border-0 shadow-lg rounded-4 p-4 p-md-5 animate-fade-up" role="alert" aria-live="polite">
              <div className="d-flex justify-content-between align-items-center mb-3">
                <span className="badge bg-danger-subtle text-danger-emphasis border border-danger-subtle">Erro 404</span>
                <small className="text-muted">Código: 404_NOT_FOUND</small>
              </div>

              <div className="text-center mb-4">
                <h1 className="display-3 fw-bold mb-0 gradient-text">404</h1>
                <p className="fs-4 fw-semibold mt-2 mb-1">Página não encontrada</p>
                <p className="text-muted mb-0">O recurso que você tentou acessar não está disponível ou pode ter sido movido.</p>
              </div>

              <div className="mb-4">
                <div className="d-flex justify-content-between align-items-center mb-1">
                  <small className="text-muted">
                    Você será redirecionado para a Home em <strong>{counter}</strong> segundo
                    {counter === 1 ? "" : "s"}...
                  </small>
                  <small className="text-muted">Redirecionando automaticamente</small>
                </div>
                <div className="progress" aria-hidden="true">
                  <div className="progress-bar" role="progressbar" style={{ width: `${progress}%` }} aria-valuenow={progress} aria-valuemin={0} aria-valuemax={100} />
                </div>
              </div>

              <div className="mb-4">
                <p className="fw-semibold mb-2">O que você pode fazer agora?</p>
                <ul className="small text-muted mb-0 ps-3">
                  <li>Voltar para a página inicial do sistema;</li>
                  <li>Retornar à página anterior;</li>
                  <li>Conferir se o endereço digitado está correto.</li>
                </ul>
              </div>

              <div className="d-flex flex-column flex-md-row gap-2">
                <button type="button" className="btn btn-primary w-100" onClick={() => navigate("/")}>
                  Ir para Home agora
                </button>

                <button type="button" className="btn btn-outline-secondary w-100" onClick={() => navigate(-1)}>
                  Voltar à página anterior
                </button>
              </div>

              <div className="mt-4 border-top pt-3 d-flex justify-content-between align-items-center flex-wrap gap-2">
                <small className="text-muted">Se o problema persistir, entre em contato com o suporte.</small>
                <small className="text-muted">
                  <span className="me-1">ID da sessão:</span>
                  <code className="small">404-{Date.now().toString(36)}</code>
                </small>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
