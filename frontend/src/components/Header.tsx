import { Link } from "react-router";
import { useState } from "react";
import logo from "../img/logo.png";
import noUser from "../img/noUser.jpg";
import searchButtom from "../img/searchButton.png";
import style from "./Header.module.scss";
import { user } from "../config/store";

interface SearchBarProps {
  search: string;
  setSearch: React.Dispatch<React.SetStateAction<string>>;
}

function Logo() {
  return (
    <Link to="/">
      <div className={style.logo}>
        <img src={logo} className="img-fluid" alt="Logo" />
      </div>
    </Link>
  );
}

function SearchBar({ search, setSearch }: SearchBarProps) {
  return (
    <div className={style.searchContainer}>
      <img src={searchButtom} className={style.searchIcon} alt="Search" />
      <input type="text" value={search} onChange={(e) => setSearch(e.target.value)} placeholder="Buscar produtos, categorias..." className={style.searchInput} />
    </div>
  );
}

function UserSide() {
  const isUserEmpty = Object.values(user).every((value) => value === "");

  if (!isUserEmpty) {
    return (
      <div className={style.userSide}>
        <img src={user.profile_img || noUser} className={style.userAvatar} alt="UserImg" />
        <p>Olá {user.first_name}</p>
      </div>
    );
  }

  return (
    <div className={style.userSide}>
      <button type="button" className={style.authButtonSecondary}>
        Registre-se
      </button>
      <button type="button" className={style.authButtonPrimary}>
        Entrar agora!
      </button>
    </div>
  );
}

function Header() {
  const [search, setSearch] = useState("");

  return (
    <header className={style.header}>
      <Logo />
      <SearchBar search={search} setSearch={setSearch} />
      <UserSide />
    </header>
  );
}

export default Header;
