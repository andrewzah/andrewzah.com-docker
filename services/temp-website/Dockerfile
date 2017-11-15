FROM anapsix/alpine-java
RUN adduser -S clojurems

COPY ./resources /srv/ms/resources
ADD target/ms-standalone.jar /srv/ms/ms.jar

USER clojurems
WORKDIR /srv/ms
CMD ["java", "-jar", "/srv/ms/ms.jar"]

