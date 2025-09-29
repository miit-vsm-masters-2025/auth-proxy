

#server {
#     listen 80;
#     server_name lms.edx.lan;
#     client_max_body_size 1000M;
#     location / {
#        proxy_pass http://127.0.0.1:8080;
#        proxy_set_header Host $host;
#        proxy_set_header X-Real-IP $remote_addr;
#        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
#        proxy_set_header X-Forwarded-Proto $scheme;
#        proxy_http_version 1.1;
#        proxy_set_header Upgrade $http_upgrade;
#        proxy_set_header Connection "Upgrade";
#    }
#}


#server {
#    listen 80;
#    server_name studio.edx.lan;
#    client_max_body_size 1000M;
#    location / {
#        proxy_pass http://127.0.0.1:8080;
#        include proxy_params;
#    }
#}
#
#server {
#    listen 80;
#    server_name apps.lms.edx.lan;
#    client_max_body_size 1000M;
#    location / {
#        proxy_pass http://127.0.0.1:8080;
#        include proxy_params;
#    }
#}
#
#server {
#    listen 80 default_server;
#    server_name _;
#    return 308 http://lms.edx.lan$request_uri;
#}


#-------------------------------------------------------------

#
#server {
#    listen 443 ssl; 
#    server_name auth.koty42.ru;
#    ssl_certificate /etc/letsencrypt/live/app.koty42.ru/fullchain.pem; # managed by Certbot
#    ssl_certificate_key /etc/letsencrypt/live/app.koty42.ru/privkey.pem; # managed by Certbot
#    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
#    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
#
#}
#
#server {
#    listen 443 ssl;
#    server_name app.koty42.ru;
#    ssl_certificate /etc/letsencrypt/live/app.koty42.ru/fullchain.pem; # managed by Certbot
#    ssl_certificate_key /etc/letsencrypt/live/app.koty42.ru/privkey.pem; # managed by Certbot
#    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
#    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
#    
#}
#
#
#
#
#server {
#    if ($host = app.koty42.ru) {
#        return 301 https://$host$request_uri;
#    } # managed by Certbot
#
#
#    listen 80;
#    server_name app.koty42.ru;
#    return 404; # managed by Certbot
#
#
#}
#
#server {
#    if ($host = auth.koty42.ru) {
#        return 301 https://$host$request_uri;
#    } # managed by Certbot
#
#
#    listen 80;
#    server_name auth.koty42.ru;
#    return 404; # managed by Certbot
#
#
#}


#----------------------------------------------------------

server {
     server_name lms.koty42.ru;
     client_max_body_size 1000M;
     location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "Upgrade";
    }
    listen 80;

    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/lms.koty42.ru/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/lms.koty42.ru/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}


server {
    server_name studio.koty42.ru;
    client_max_body_size 1000M;
    location / {
        proxy_pass http://127.0.0.1:8080;
        include proxy_params;
    }
    listen 80;

    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/lms.koty42.ru/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/lms.koty42.ru/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}

server {
    server_name apps.lms.koty42.ru;
    client_max_body_size 1000M;
    location / {
        proxy_pass http://127.0.0.1:8080;
        include proxy_params;
    }
    listen 80;
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/lms.koty42.ru/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/lms.koty42.ru/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}



#
#server {
#    if ($host = lms.koty42.ru) {
#        return 301 https://$host$request_uri;
#    } # managed by Certbot
#
#
#     listen 80;
#     server_name lms.koty42.ru;
#    return 404; # managed by Certbot
#
#
#}
#
#server {
#    if ($host = apps.lms.koty42.ru) {
#        return 301 https://$host$request_uri;
#    } # managed by Certbot
#
#
#    listen 80;
#    server_name apps.lms.koty42.ru;
#    return 404; # managed by Certbot
#
#
#}
#
#
#server {
#    if ($host = studio.koty42.ru) {
#        return 301 https://$host$request_uri;
#    } # managed by Certbot
#
#
#    listen 80;
#    server_name studio.koty42.ru;
#    return 404; # managed by Certbot
#
#
#}
