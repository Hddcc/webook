import Head from "next/head";
import Link from "next/link";

import styles from "@/styles/Signup.module.css";

export default function LoginPlaceholder() {
  return (
    <>
      <Head>
        <title>登录 | 小微书</title>
      </Head>
      <main className={styles.page}>
        <section className={styles.panel}>
          <div className={styles.heading}>
            <span className={styles.eyebrow}>WEBOOK</span>
            <h1>登录页暂未创建</h1>
            <p>当前练习仅包含注册页面。</p>
          </div>
          <Link className={styles.backLink} href="/users/signup">
            返回注册
          </Link>
        </section>
      </main>
    </>
  );
}
