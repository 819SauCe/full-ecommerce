import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import AuthLayout from "../layout/AuthLayout";
import TextInput from "../components/TextInput";
import PasswordInput from "../components/PasswordInput";
import Checkbox from "../components/Checkbox";
import Button from "../components/Button";
import FormError from "../components/FormError";

import { loginRequest } from "../services/authService";

const LoginPage: React.FC = () => {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [rememberMe, setRememberMe] = useState(true);
  const [formError, setFormError] = useState<string | undefined>();
  const [fieldErrors, setFieldErrors] = useState<{
    email?: string;
    password?: string;
  }>({});
  const [loading, setLoading] = useState(false);

  function validate() {
    const errors: { email?: string; password?: string } = {};

    if (!email) errors.email = "Informe seu e-mail.";
    else if (!/\S+@\S+\.\S+/.test(email)) errors.email = "E-mail inválido.";

    if (!password) errors.password = "Informe sua senha.";
    else if (password.length < 6) errors.password = "A senha deve ter ao menos 6 caracteres.";

    setFieldErrors(errors);

    return Object.keys(errors).length === 0;
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setFormError(undefined);

    if (!validate()) return;

    try {
      setLoading(true);
      await loginRequest({ email, password });
      navigate("/");
    } catch (err: any) {
      setFormError(err.message || "Erro ao fazer login. Tente novamente.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <AuthLayout title="Entrar na sua conta" subtitle="Acesse para acompanhar pedidos, favoritos e muito mais.">
      <FormError message={formError} />

      <form onSubmit={handleSubmit} noValidate>
        <TextInput label="E-mail" name="email" type="email" placeholder="seuemail@exemplo.com" value={email} onChange={(e) => setEmail(e.target.value)} error={fieldErrors.email} />

        <PasswordInput label="Senha" name="password" placeholder="Sua senha" value={password} onChange={(e) => setPassword(e.target.value)} error={fieldErrors.password} />

        <div className="d-flex justify-content-between align-items-center mb-3">
          <Checkbox label="Manter conectado" name="rememberMe" checked={rememberMe} onChange={(e) => setRememberMe(e.target.checked)} />

          <button type="button" className="btn btn-link p-0" onClick={() => navigate("/forgot-password")}>
            Esqueci minha senha
          </button>
        </div>

        <Button type="submit" variant="primary" size="lg" fullWidth loading={loading}>
          Entrar
        </Button>
      </form>

      <hr className="my-4" />

      <p className="text-center mb-0" style={{ fontSize: 14 }}>
        Ainda não tem conta?{" "}
        <button type="button" className="btn btn-link p-0" onClick={() => navigate("/register")}>
          Criar conta
        </button>
      </p>
    </AuthLayout>
  );
};

export default LoginPage;
