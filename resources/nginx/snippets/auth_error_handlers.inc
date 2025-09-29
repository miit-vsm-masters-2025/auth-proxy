location @login {
    return 302 https://auth.koty42.ru/login?return_to=$scheme://$host$request_uri;
}

location @forbidden {
    return 403;
}