FROM python:3

WORKDIR /usr/src/app

RUN pip install --no-cache-dir "jsonrpcclient[requests]"

COPY . .

CMD [ "python", "./main.py" ]