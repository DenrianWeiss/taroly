FROM ubuntu:latest

RUN apt update && apt install -y curl git && curl -L https://foundry.paradigm.xyz | bash && ~/.foundry/bin/foundryup
ENV PATH="/root/.foundry/bin:${PATH}"

COPY taroly /usr/local/bin/taroly
RUN chmod +x /usr/local/bin/taroly

ENTRYPOINT ["/usr/local/bin/taroly"]
