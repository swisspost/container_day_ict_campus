FROM python:3.13
WORKDIR /code
ENV FLASK_APP=hello.py
RUN pip install flask
EXPOSE 8000
COPY . .
CMD ["flask", "run", "--port=8000","--host=0.0.0.0"]
