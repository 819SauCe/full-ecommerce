import logo from "../img/logo.png";
import noUser from "../img/noUser.jpg";
import searchButtom from "../img/searchButton.png";
import style from "./Header.module.scss";
import { user } from "../config/store";

const isUserEmpty = Object.values(user).every((value) => value === "");

function Logo() {
  return (
    <div className={style.logo}>
      <img src={logo} className="img-fluid" alt="Logo" />
    </div>
  );
}

function SearchBar() {
  return (
    <div className={style.searchContainer}>
      <img src={searchButtom} className={style.searchIcon} alt="Search" />
      <input type="text" placeholder="Buscar produtos, categorias..." className={style.searchInput} />
    </div>
  );
}

function UserSide() {
  if (!isUserEmpty) {
    return (
      <div className={style.userSide}>
        <img src={user.profile_img || noUser} className={style.userAvatar} alt="UserImg" />
        <p>Olá {user.first_name}</p>
      </div>
    );
  } else {
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
}

function Header() {
  return (
    <header className={style.header}>
      <Logo />
      <SearchBar />
      <UserSide />
    </header>
  );
}

export default Header;
