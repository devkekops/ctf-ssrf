# ctf-ssrf

docker build -t ctf-ssrf .

docker run -e SERVER_ADDRESS='0.0.0.0:80' -e FLAG='...' -p 80:80 ctf-ssrf