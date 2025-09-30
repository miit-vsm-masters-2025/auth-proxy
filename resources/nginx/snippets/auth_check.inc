# Проверка авторизации
auth_request /_auth;

# Если сервис проверки авторизации ответил 401/403 (необходимо добавить auth_error_handlers.inc для коректной работы)
auth_request_set $auth_set_cookie  $upstream_http_set_cookie;
# auth_request_set $auth_set_header  $upstream_http_set_header;
error_page 401 = @login;
error_page 403 = @forbidden;

# Забираем заголовки из ответа auth-сервиса и прокидываем в бэкенд
auth_request_set $user_id     $upstream_http_x_user_id;
# auth_request_set $user_roles  $upstream_http_x_user_roles;
auth_request_set $auth_status $upstream_status;

# (Для диагностики) показать статус auth-подзапроса
add_header X-Auth-Status $auth_status always;

# Пробрасываем id пользователя
proxy_set_header X-User-Id    $user_id;

# Стандартные заголовки
proxy_set_header Host              $host;
proxy_set_header X-Forwarded-Proto $scheme;
proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
proxy_set_header X-Real-IP         $remote_addr;
