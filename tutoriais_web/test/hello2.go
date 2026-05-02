package main

import (
	"fmt"
	"net/http"
)

const blogPage = `<!doctype html>
<html lang="pt-BR">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Blog em Go</title>
  <style>
    :root {
      color-scheme: dark;
      font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      background: #101418;
      color: #f3f7f4;
    }

    * {
      box-sizing: border-box;
    }

    body {
      min-height: 100vh;
      margin: 0;
      overflow-x: hidden;
      background:
        radial-gradient(circle at 18% 18%, rgba(54, 181, 138, 0.24), transparent 26rem),
        radial-gradient(circle at 82% 12%, rgba(120, 168, 255, 0.18), transparent 24rem),
        linear-gradient(145deg, #101418 0%, #18201d 52%, #0f1518 100%);
    }

    main {
      min-height: 100vh;
      display: grid;
      place-items: center;
      padding: 40px 20px;
      transform: translateY(70px) scale(0.98);
      opacity: 0;
      animation: settle 1100ms cubic-bezier(.2, .8, .2, 1) forwards;
    }

    .panel {
      width: min(880px, 100%);
      border: 1px solid rgba(255, 255, 255, 0.14);
      border-radius: 8px;
      padding: clamp(24px, 6vw, 64px);
      background: rgba(16, 20, 24, 0.72);
      box-shadow: 0 28px 90px rgba(0, 0, 0, 0.34);
      backdrop-filter: blur(18px);
    }

    .eyebrow {
      margin: 0 0 16px;
      color: #7ddfbb;
      font-size: 0.88rem;
      font-weight: 700;
      letter-spacing: 0;
      text-transform: uppercase;
    }

    h1 {
      max-width: 760px;
      margin: 0;
      font-size: clamp(2.3rem, 8vw, 5.8rem);
      line-height: 0.96;
      letter-spacing: 0;
    }

    .lead {
      max-width: 660px;
      margin: 24px 0 0;
      color: #c9d6d0;
      font-size: clamp(1rem, 2vw, 1.22rem);
      line-height: 1.7;
    }

    .messages {
      display: grid;
      gap: 12px;
      margin-top: 34px;
    }

    .message {
      width: fit-content;
      max-width: 100%;
      padding: 12px 14px;
      border: 1px solid rgba(255, 255, 255, 0.12);
      border-radius: 8px;
      background: rgba(255, 255, 255, 0.08);
      color: #eef7f2;
      opacity: 0;
      transform: translateY(18px);
      animation: reveal 650ms ease forwards;
    }

    .message:nth-child(1) {
      animation-delay: 650ms;
    }

    .message:nth-child(2) {
      animation-delay: 950ms;
    }

    .message:nth-child(3) {
      animation-delay: 1250ms;
    }

    .actions {
      display: flex;
      flex-wrap: wrap;
      gap: 12px;
      margin-top: 34px;
    }

    button,
    a {
      min-height: 44px;
      border: 0;
      border-radius: 8px;
      padding: 0 18px;
      font: inherit;
      font-weight: 700;
      cursor: pointer;
    }

    button {
      background: #7ddfbb;
      color: #0d1714;
    }

    a {
      display: inline-flex;
      align-items: center;
      color: #f3f7f4;
      text-decoration: none;
      background: rgba(255, 255, 255, 0.1);
    }

    .pulse {
      animation: pulse 600ms ease;
    }

    @keyframes settle {
      0% {
        opacity: 0;
        transform: translateY(70px) scale(0.98);
      }

      100% {
        opacity: 1;
        transform: translateY(0) scale(1);
      }
    }

    @keyframes reveal {
      to {
        opacity: 1;
        transform: translateY(0);
      }
    }

    @keyframes pulse {
      50% {
        transform: translateY(-6px);
      }
    }

    @media (max-width: 640px) {
      .panel {
        padding: 24px;
      }

      h1 {
        font-size: 3rem;
      }
    }
  </style>
</head>
<body>
  <main id="stage">
    <section class="panel">
      <p class="eyebrow">Experimento em Go</p>
      <h1>Um blog que acorda em movimento.</h1>
      <p class="lead">
        Esta pagina e so um teste rapido: o servidor Go entrega o HTML, a tela sobe ao entrar,
        e as mensagens aparecem como se o site estivesse se ajeitando para receber o leitor.
      </p>

      <div class="messages" aria-live="polite">
        <div class="message">Carregando postagens em Markdown...</div>
        <div class="message">Separando jogos em JavaScript...</div>
        <div class="message">Preparando curiosidades para a pagina inicial...</div>
      </div>

      <div class="actions">
        <button id="again" type="button">Repetir efeito</button>
        <a href="/health">Ver status</a>
      </div>
    </section>
  </main>

  <script>
    const stage = document.querySelector("#stage");
    const button = document.querySelector("#again");

    button.addEventListener("click", () => {
      stage.classList.remove("pulse");
      void stage.offsetWidth;
      stage.classList.add("pulse");
    });
  </script>
</body>
</html>`

func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, blogPage)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "ok")
}
