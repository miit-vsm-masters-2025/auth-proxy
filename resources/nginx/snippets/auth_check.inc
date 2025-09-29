# 1) Проверка авторизации
auth_request /_auth;

# 2) Забираем заголовки из ответа auth-сервиса и прокидываем в бэкенд
auth_request_set $user_id     $upstream_http_x_user_id;
auth_request_set $user_roles  $upstream_http_x_user_roles;
auth_request_set $auth_status $upstream_status;

# На 401/403 перенаправляем/отвечаем отдельно (необходимо добавить auth_error_handlers.inc)
error_page 401 = @login;
error_page 403 = @forbidden;

# (Для диагностики) показать статус auth-подзапроса
add_header X-Auth-Status $auth_status always;

# Пробрасываем идентификацию пользователя в бэкенд
proxy_set_header X-User-Id    $user_id;
proxy_set_header X-User-Roles $user_roles;

# Стандартные заголовки
proxy_set_header Host              $host;
proxy_set_header X-Forwarded-Proto $scheme;
proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
proxy_set_header X-Real-IP         $remote_addr;
