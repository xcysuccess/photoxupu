FROM csighub.tencentyun.com/medipedia/medi-saas-go:latest

RUN mkdir -p /app/logs/
ADD sql-operator /usr/local/services/sql-operator/
COPY script/supervisord.ini /etc/supervisord.d/
COPY script/kick_start.sh /etc/kickStart.d/

# 创建存放导入文件的目录
RUN mkdir -p /usr/local/services/sql-operator/importfile
# fix a protocol buffer namespace conflict
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT warn

# 修改镜像的显示字符集
ENV LANG en_US.UTF-8