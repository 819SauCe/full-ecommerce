import React, { useEffect, useState, useRef } from "react";
import styles from "./Banner.module.scss";
import { base_api, store_banner } from "../config/api";

interface BannerType {
  id: number;
  img: string;
  title?: string;
  subtitle?: string;
  button_to?: string;
  button_label?: string;
}

export default function Banner() {
  const [banners, setBanners] = useState<BannerType[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [currentIndex, setCurrentIndex] = useState(0);

  const intervalRef = useRef<number | null>(null);

  useEffect(() => {
    const load = async () => {
      try {
        const res = await fetch(`${base_api}${store_banner}`);
        if (!res.ok) throw new Error("Erro ao buscar banners");
        const data = await res.json();
        setBanners(data);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };
    load();
  }, []);

  useEffect(() => {
    if (!banners.length) return;
    if (intervalRef.current) clearInterval(intervalRef.current);
    intervalRef.current = window.setInterval(() => {
      setCurrentIndex((prev) => (prev + 1) % banners.length);
    }, 5000);
    return () => {
      if (intervalRef.current) clearInterval(intervalRef.current);
    };
  }, [banners]);

  const next = () => setCurrentIndex((prev) => (prev + 1) % banners.length);
  const prev = () => setCurrentIndex((prev) => (prev === 0 ? banners.length - 1 : prev - 1));

  if (loading) return <div>Carregando...</div>;
  if (error) return <div>Erro: {error}</div>;
  if (!banners.length) return <div>Nenhum banner encontrado</div>;

  const banner = banners[currentIndex];

  return (
    <div className={styles.carousel}>
      <button className={`${styles.arrow} ${styles.arrowLeft}`} onClick={prev}>
        ‹
      </button>

      <div className={styles.slide}>
        <img src={banner.img} className={styles.image} />

        <div className={styles.content}>
          {banner.title && <h2 className={styles.title}>{banner.title}</h2>}
          {banner.subtitle && <p className={styles.subtitle}>{banner.subtitle}</p>}
          {banner.button_to && (
            <a href={banner.button_to} className={styles.button}>
              {banner.button_label}
            </a>
          )}
        </div>
      </div>

      <button className={`${styles.arrow} ${styles.arrowRight}`} onClick={next}>
        ›
      </button>

      <div className={styles.dots}>
        {banners.map((_, i) => (
          <button key={i} className={`${styles.dot} ${i === currentIndex ? styles.active : ""}`} onClick={() => setCurrentIndex(i)} />
        ))}
      </div>
    </div>
  );
}
