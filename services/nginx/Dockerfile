FROM nginx

LABEL maintainer="Dinda <dindasigma@gmail.com>"

COPY nginx.conf /etc/nginx/nginx.conf

# access the server via localhost:8080 on our host.
EXPOSE 8080

# Start Nginx when the container has provisioned.
CMD ["nginx", "-g", "daemon off;"]