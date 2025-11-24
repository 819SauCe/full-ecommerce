import { Link, useLocation } from "react-router";
import style from "./NaBar.module.scss";

const navLinks = [
  { to: "/", label: "Início" },
  { to: "/store", label: "Loja" },
  {
    to: "/categories",
    label: "Categorias",
    children: [
      { to: "/categories/celulares", label: "Celulares" },
      { to: "/categories/notebooks", label: "Notebooks" },
      { to: "/categories/consoles", label: "Consoles" },
      { to: "/categories/pc-gamer", label: "PC Gamer" },
    ],
  },
  { to: "/offers", label: "Ofertas" },
  { to: "/support", label: "Suporte" },
];

function NavBar() {
  const location = useLocation();

  return (
    <nav className={style.navbar}>
      <ul className={style.navList}>
        {navLinks.map((link) => {
          const isActive = location.pathname === link.to || (link.to !== "/" && location.pathname.startsWith(link.to));

          const hasDropdown = Array.isArray(link.children);

          return (
            <li key={link.to} className={`${style.navItem} ${hasDropdown ? style.navItemHasDropdown : ""}`}>
              <Link to={link.to} className={`${style.navLink} ${isActive ? style.navLinkActive : ""}`}>
                {link.label}
              </Link>

              {hasDropdown && (
                <ul className={style.dropdown}>
                  {link.children!.map((child) => (
                    <li key={child.to} className={style.dropdownItem}>
                      <Link to={child.to} className={style.dropdownLink}>
                        {child.label}
                      </Link>
                    </li>
                  ))}
                </ul>
              )}
            </li>
          );
        })}
      </ul>
    </nav>
  );
}

export default NavBar;
