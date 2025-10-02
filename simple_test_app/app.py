#!/usr/bin/env python3
"""
Simple header echo server.
Listens on port 8181 and returns an HTML page with all incoming request headers.

Usage:
python3 /app/header_echo.py
# Then open http://localhost:8181 in your browser
# Or: curl -i -H 'X-Test: 123' http://localhost:8181/
"""
from http.server import ThreadingHTTPServer, BaseHTTPRequestHandler
import html
import sys

class HeaderEchoHandler(BaseHTTPRequestHandler):
  server_version = "HeaderEcho/1.0"

  def _respond_html(self):
      # Build rows for headers table
      rows = []
      for key, value in self.headers.items():
          rows.append(
              f"<tr><td class='k'>{html.escape(key)}</td>"
              f"<td class='v'>{html.escape(value)}</td></tr>"
          )

      rows_html = "\n".join(rows) if rows else "<tr><td colspan='2' class='empty'>нет заголовков</td></tr>"

      body = f"""<!doctype html>
<html lang="ru">
<head>
<meta charset="utf-8">
<title>Header Echo</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
  body {{ font-family: system-ui, -apple-system, Segoe UI, Roboto, sans-serif; margin: 2rem; }}
  h1 {{ margin: 0 0 0.5rem; font-size: 1.4rem; }}
  .meta {{ color: #555; margin-bottom: 1rem; }}
  table {{ border-collapse: collapse; width: 100%; max-width: 960px; }}
  th, td {{ border: 1px solid #ddd; padding: 8px; vertical-align: top; }}
  th {{ background: #f6f6f6; text-align: left; }}
  td.k {{ white-space: nowrap; font-weight: 600; width: 1%; }}
  td.v {{ word-break: break-all; }}
  .empty {{ text-align: center; color: #777; }}
  footer {{ margin-top: 1rem; color: #777; font-size: 0.9rem; }}
  code {{ background: #f6f6f6; padding: 0 4px; border-radius: 3px; }}
</style>
</head>
<body>
<h1>Полученные HTTP-заголовки</h1>
<div class="meta">
  Метод: <code>{html.escape(self.command)}</code>,
  Путь: <code>{html.escape(self.path)}</code>,
  Клиент: <code>{html.escape(self.client_address[0])}</code>
</div>
<table>
  <thead>
    <tr><th>Заголовок</th><th>Значение</th></tr>
  </thead>
  <tbody>
    {rows_html}
  </tbody>
</table>
<footer>HeaderEcho/1.0 — Python http.server</footer>
</body>
</html>"""

      data = body.encode("utf-8")
      self.send_response(200)
      self.send_header("Content-Type", "text/html; charset=utf-8")
      self.send_header("Content-Length", str(len(data)))
      self.end_headers()

      # Do not write body for HEAD
      if self.command != "HEAD":
          self.wfile.write(data)

  # Support common methods
  def do_GET(self): self._respond_html()
  def do_HEAD(self): self._respond_html()
  def do_POST(self): self._respond_html()
  def do_PUT(self): self._respond_html()
  def do_DELETE(self): self._respond_html()
  def do_PATCH(self): self._respond_html()

  # Cleaner access log
  def log_message(self, fmt, *args):
      sys.stderr.write("%s - - [%s] %s\n" % (
          self.client_address[0],
          self.log_date_time_string(),
          fmt % args
      ))

def run(host: str = "127.0.0.1", port: int = 8181):
  with ThreadingHTTPServer((host, port), HeaderEchoHandler) as httpd:
      print(f"Serving on http://{host}:{port}", flush=True)
      try:
          httpd.serve_forever()
      except KeyboardInterrupt:
          print("\nShutting down...", flush=True)

if __name__ == "__main__":
  run()