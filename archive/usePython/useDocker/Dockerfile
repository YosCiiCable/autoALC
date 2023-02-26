FROM python:3
WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

#RUN	export DISPLAY=:0; \
#	export PULSE_SERVER=/mnt/wslg/PulseServer; \
#	export WAYLAND_DISPLAY=wayland-0; \
#	export XDG_RUNTIME_DIR=/run/user/1000/ ;

#CMD [ "python", "./InputAuto.py" ]
