FROM ubuntu:latest

COPY taroly /usr/local/bin/taroly
RUN chmod +x /usr/local/bin/taroly
RUN apt update && apt install -y curl && curl -L https://foundry.paradigm.xyz | bash && ~/.foundry/bin/foundryup \
ENV PATH="/root/.foundry/bin:${PATH}"
ENTRYPOINT ["/usr/local/bin/taroly"]
