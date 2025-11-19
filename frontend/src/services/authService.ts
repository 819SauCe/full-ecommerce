import { API_ROUTES } from "../config/apiConfig";

export interface LoginPayload {
  email: string;
  password: string;
}

export interface RegisterPayload {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}

export async function loginRequest(data: LoginPayload) {
  const response = await fetch(API_ROUTES.login, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    const errText = await response.text();
    throw new Error(errText || "Erro ao fazer login");
  }

  return response.json().catch(() => ({}));
}

export async function registerRequest(data: RegisterPayload) {
  const response = await fetch(API_ROUTES.register, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      first_name: data.firstName,
      last_name: data.lastName,
      email: data.email,
      password: data.password,
    }),
  });

  if (!response.ok) {
    const errText = await response.text();
    throw new Error(errText || "Erro ao criar conta");
  }
  try {
    return await response.json();
  } catch {
    return {};
  }
}
