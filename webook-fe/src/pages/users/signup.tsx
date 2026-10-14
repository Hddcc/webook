import axios from "axios";
import { Eye, EyeOff } from "lucide-react";
import Head from "next/head";
import Link from "next/link";
import { FormEvent, useState } from "react";

import { http } from "@/axios/http";
import styles from "@/styles/Signup.module.css";

type Feedback = {
  kind: "success" | "error";
  message: string;
};

export default function SignUpPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [feedback, setFeedback] = useState<Feedback | null>(null);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setFeedback(null);

    if (password !== confirmPassword) {
      setFeedback({ kind: "error", message: "两次输入的密码不一致" });
      return;
    }

    setSubmitting(true);
    try {
      const response = await http.post("/users", {
        email,
        password,
        confirmPassword,
      });
      setFeedback({
        kind: "success",
        message: response.data?.message ?? "注册请求已提交",
      });
    } catch (error) {
      let message = "请求失败，请确认后端已在 8085 端口启动";
      if (axios.isAxiosError(error)) {
        message = error.response?.data?.message ?? error.message ?? message;
      }
      setFeedback({ kind: "error", message });
    } finally {
      setSubmitting(false);
    }
  }

  return (
    <>
      <Head>
        <title>注册 | 小微书</title>
        <meta name="description" content="小微书用户注册" />
      </Head>
      <main className={styles.page}>
        <section className={styles.panel} aria-labelledby="signup-title">
          <div className={styles.heading}>
            <span className={styles.eyebrow}>WEBOOK</span>
            <h1 id="signup-title">创建账号</h1>
            <p>注册后开始记录与分享。</p>
          </div>

          <form className={styles.form} onSubmit={handleSubmit}>
            <label className={styles.field}>
              <span>邮箱</span>
              <input
                type="email"
                name="email"
                autoComplete="email"
                placeholder="name@example.com"
                value={email}
                onChange={(event) => setEmail(event.target.value)}
                required
              />
            </label>

            <label className={styles.field}>
              <span>密码</span>
              <div className={styles.passwordControl}>
                <input
                  type={showPassword ? "text" : "password"}
                  name="password"
                  autoComplete="new-password"
                  placeholder="请输入密码"
                  value={password}
                  onChange={(event) => setPassword(event.target.value)}
                  minLength={6}
                  required
                />
                <button
                  type="button"
                  className={styles.visibilityButton}
                  onClick={() => setShowPassword((visible) => !visible)}
                  aria-label={showPassword ? "隐藏密码" : "显示密码"}
                  title={showPassword ? "隐藏密码" : "显示密码"}
                >
                  {showPassword ? <EyeOff aria-hidden="true" /> : <Eye aria-hidden="true" />}
                </button>
              </div>
            </label>

            <label className={styles.field}>
              <span>确认密码</span>
              <div className={styles.passwordControl}>
                <input
                  type={showConfirmPassword ? "text" : "password"}
                  name="confirmPassword"
                  autoComplete="new-password"
                  placeholder="请再次输入密码"
                  value={confirmPassword}
                  onChange={(event) => setConfirmPassword(event.target.value)}
                  minLength={6}
                  required
                />
                <button
                  type="button"
                  className={styles.visibilityButton}
                  onClick={() => setShowConfirmPassword((visible) => !visible)}
                  aria-label={showConfirmPassword ? "隐藏确认密码" : "显示确认密码"}
                  title={showConfirmPassword ? "隐藏确认密码" : "显示确认密码"}
                >
                  {showConfirmPassword ? <EyeOff aria-hidden="true" /> : <Eye aria-hidden="true" />}
                </button>
              </div>
            </label>

            {feedback && (
              <p className={styles[feedback.kind]} role="status">
                {feedback.message}
              </p>
            )}

            <div className={styles.actions}>
              <button className={styles.submitButton} type="submit" disabled={submitting}>
                {submitting ? "提交中..." : "注册"}
              </button>
              <span>
                已有账号？<Link href="/users/login">登录</Link>
              </span>
            </div>
          </form>
        </section>
      </main>
    </>
  );
}
