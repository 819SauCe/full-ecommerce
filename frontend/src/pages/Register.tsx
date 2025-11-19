import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import AuthLayout from "../layout/AuthLayout";
import TextInput from "../components/TextInput";
import PasswordInput from "../components/PasswordInput";
import Checkbox from "../components/Checkbox";
import Button from "../components/Button";
import FormError from "../components/FormError";
import { registerRequest } from "../services/authService";

const RegisterPage: React.FC = () => {
  const navigate = useNavigate();

  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passwordConfirm, setPasswordConfirm] = useState("");
  const [acceptTerms, setAcceptTerms] = useState(false);

  const [formError, setFormError] = useState<string | undefined>();
  const [fieldErrors, setFieldErrors] = useState<{
    name?: string;
    email?: string;
    password?: string;
    passwordConfirm?: string;
    acceptTerms?: string;
  }>({});
  const [loading, setLoading] = useState(false);

  function validate() {
    const errors: typeof fieldErrors = {};
    const trimmedName = name.trim();

    if (!trimmedName) {
      errors.name = "Informe seu nome completo.";
    } else if (trimmedName.length < 3) {
      errors.name = "O nome deve ter ao menos 3 caracteres.";
    } else if (trimmedName.split(/\s+/).length < 2) {
      errors.name = "Informe nome e sobrenome.";
    }

    if (!email) {
      errors.email = "Informe seu e-mail.";
    } else if (!/\S+@\S+\.\S+/.test(email)) {
      errors.email = "E-mail inválido.";
    }

    if (!password) {
      errors.password = "Informe uma senha.";
    } else if (password.length < 6) {
      errors.password = "A senha deve ter ao menos 6 caracteres.";
    }

    if (!passwordConfirm) {
      errors.passwordConfirm = "Confirme sua senha.";
    } else if (passwordConfirm !== password) {
      errors.passwordConfirm = "As senhas não coincidem.";
    }

    if (!acceptTerms) {
      errors.acceptTerms = "Você precisa aceitar os termos para continuar.";
    }

    setFieldErrors(errors);
    return Object.keys(errors).length === 0;
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setFormError(undefined);

    if (!validate()) return;

    try {
      setLoading(true);

      const trimmedName = name.trim();
      const [firstName, ...rest] = trimmedName.split(/\s+/);
      const lastName = rest.join(" ") || firstName;

      await registerRequest({
        firstName,
        lastName,
        email,
        password,
      });
      navigate("/login");
    } catch (err: any) {
      setFormError(err.message || "Erro ao criar conta. Tente novamente em instantes.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <AuthLayout title="Criar sua conta" subtitle="Cadastre-se para acompanhar pedidos, favoritos e ofertas exclusivas.">
      <FormError message={formError} />

      <form onSubmit={handleSubmit} noValidate>
        <TextInput label="Nome completo" name="name" placeholder="Seu nome" value={name} onChange={(e) => setName(e.target.value)} error={fieldErrors.name} />

        <TextInput label="E-mail" name="email" type="email" placeholder="seuemail@exemplo.com" value={email} onChange={(e) => setEmail(e.target.value)} error={fieldErrors.email} />

        <PasswordInput label="Senha" name="password" placeholder="Crie uma senha" value={password} onChange={(e) => setPassword(e.target.value)} error={fieldErrors.password} />

        <PasswordInput label="Confirmar senha" name="passwordConfirm" placeholder="Repita a senha" value={passwordConfirm} onChange={(e) => setPasswordConfirm(e.target.value)} error={fieldErrors.passwordConfirm} />

        <div className="mb-3">
          <Checkbox label="Li e aceito os Termos de Uso e Política de Privacidade" name="acceptTerms" checked={acceptTerms} onChange={(e) => setAcceptTerms(e.target.checked)} />
          {fieldErrors.acceptTerms && (
            <div className="text-danger mt-1" style={{ fontSize: "0.85rem" }}>
              {fieldErrors.acceptTerms}
            </div>
          )}
        </div>

        <Button type="submit" variant="primary" size="lg" fullWidth loading={loading}>
          Criar conta
        </Button>
      </form>

      <hr className="my-4" />

      <p className="text-center mb-0" style={{ fontSize: 14 }}>
        Já tem uma conta?{" "}
        <button type="button" className="btn btn-link p-0" style={{ fontSize: 14 }} onClick={() => navigate("/login")}>
          Entrar
        </button>
      </p>
    </AuthLayout>
  );
};

export default RegisterPage;
