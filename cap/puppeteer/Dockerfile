FROM cap.dudaji.com:31480/aiops/node:12

RUN apt-get update 

RUN apt-get install -y vim libgtk2.0-0 libgtk-3-0 libnotify-dev 
RUN apt-get install -y libgconf-2-4 libnss3 libxss1 
RUN apt-get install -y libasound2 libxtst6 xauth xvfb
RUN apt-get install -y libgbm-dev

WORKDIR /app

COPY . /app

RUN yarn install

RUN yarn add puppeteer


CMD ["node", "ui_test.js","TARGET_SERVER"]