From python:3.7.5

RUN pip install --upgrade pip
RUN pip install numpy
RUN pip install matplotlib
RUN pip install pandas
RUN pip install tabulate
RUN pip install flask

WORKDIR /TimeSheetsApp/src/work
Add . /TimeSheetsApp/src/work
CMD ["flask", "run","--host", "0.0.0.0", "--port", "5000"]