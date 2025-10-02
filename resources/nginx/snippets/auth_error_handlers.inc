location @login {
    # всегда инвалидируем куки при ошибке авторизации
    add_header Set-Cookie $auth_set_cookie always;
    # add_header Set-Header $auth_set_header always;
    return 302 https://auth.koty42.ru/signin?return_to=$scheme://$host$request_uri;
}

location @forbidden {
    # 403 пока просто проксируем тк ролей нет и не ясно зачем оно тут нужно
    return 403;
}