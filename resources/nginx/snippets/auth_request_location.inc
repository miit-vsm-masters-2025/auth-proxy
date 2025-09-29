# Внутренняя локация для подзапроса авторизации
internal;

proxy_pass http://auth_service/auth;
proxy_http_version 1.1;

proxy_pass_request_body off;
proxy_set_header Content-Length "";

# Пробрасываем исходные атрибуты запроса (auth-сервис может их учитывать)
proxy_set_header X-Original-Method $request_method;
proxy_set_header X-Original-URI    $request_uri;
proxy_set_header X-Original-Host   $host;
proxy_set_header X-Forwarded-Proto $scheme;
proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
proxy_set_header X-Real-IP         $remote_addr;

# Передаём учётные данные (если используются)
proxy_set_header Cookie        $http_cookie;
proxy_set_header Authorization $http_authorization;

# Своё (более строгое) время ожидания для auth
proxy_connect_timeout 2s;
proxy_send_timeout    2s;
proxy_read_timeout    2s;