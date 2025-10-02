package utils

import (
	"fmt"
	"html"
	"strings"
)

// html templates are  better choice, but to make it simple - avoid outer files
func RenderSignInHTML(redirect, login, errMsg string) string {
	escRedirect := html.EscapeString(redirect)
	escLogin := html.EscapeString(login)
	errBlock := ""
	if strings.TrimSpace(errMsg) != "" {
		errBlock = fmt.Sprintf(`<div class="error">%s</div>`, html.EscapeString(errMsg))
	}

	return fmt.Sprintf(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Sign in</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="Content-Security-Policy" content="default-src 'none'; style-src 'self' 'unsafe-inline'; form-action 'self'; base-uri 'none'">
  <style>
    :root { color-scheme: light dark; }
    body { font-family: system-ui, -apple-system, Segoe UI, Roboto, sans-serif; margin: 0; padding: 0; }
    .wrap { display: grid; place-items: center; min-height: 100vh; padding: 16px; }
    .card {
      width: 100%%; max-width: 380px; padding: 24px; border-radius: 12px;
      border: 1px solid rgba(128,128,128,.25);
      box-shadow: 0 6px 18px rgba(0,0,0,.08);
      background: canvas;
    }
    h1 { margin: 0 0 16px; font-size: 20px; }
    .field { margin: 12px 0; }
    label { display: block; font-size: 14px; margin-bottom: 6px; }
    input[type="text"], input[type="password"] {
      width: 100%%; padding: 10px 12px; border-radius: 8px;
      border: 1px solid rgba(128,128,128,.35); background: canvas; color: canvastext;
    }
    .error {
      margin: 8px 0 12px; padding: 10px 12px; border-radius: 8px;
      background: #fee; color: #a00; border: 1px solid #f99;
    }
    button {
      width: 100%%; padding: 10px 14px; border: 0; border-radius: 8px;
      background: #2563eb; color: #fff; font-weight: 600; cursor: pointer;
    }
    .hint { margin-top: 10px; font-size: 12px; opacity: .8; }
  </style>
</head>
<body>
  <div class="wrap">
    <div class="card">
      <h1>Sign in</h1>
      %s
      <form method="post" action="/user/login">
        <input type="hidden" name="return_to" value="%s">
        <div class="field">
          <label for="login">Login</label>
          <input id="login" name="login" type="text" autocomplete="username" value="%s" required>
        </div>
        <div class="field">
          <label for="password">Password</label>
          <input id="password" name="password" type="password" autocomplete="current-password" required>
        </div>
        <button type="submit">Sign in</button>
        <div class="hint">Demo: admin / password</div>
      </form>
    </div>
  </div>
</body>
</html>`, errBlock, escRedirect, escLogin)
}
