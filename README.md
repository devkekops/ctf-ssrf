# ctf-ssrf

docker build -t ctf-ssrf .

# use unpredictable port 
docker run -e SERVER_ADDRESS='0.0.0.0:<port>' -e FLAG='...' -p <port>:80 ctf-ssrf