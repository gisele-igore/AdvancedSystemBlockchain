FROM node:8.12.0

WORKDIR /home/node

COPY package.json /home/node

RUN npm install

ENV WIMTK_VERSION v0.6.0
RUN wget https://github.com/wid/wimtk/releases/download/${WIMTK_VERSION}/wimtk-alpine -O /usr/local/bin/wimtk && chmod 755 /usr/local/bin/wimtk

ADD . /home/node

RUN npm run build; exit 0;

CMD npm run start
#CMD npm run start:prod
