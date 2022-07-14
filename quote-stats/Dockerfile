FROM python:3.9.1-slim 

LABEL org.opencontainers.image.source https://github.com/s1ntaxe770r/gopher-quotes


WORKDIR /app 

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

COPY requirements.txt .

RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt 

COPY . .

EXPOSE  8080
# start server with 4 workers 

CMD ["gunicorn","-w", "4", "--log-level","info", "main:app", "-b", "0.0.0.0:8081"]