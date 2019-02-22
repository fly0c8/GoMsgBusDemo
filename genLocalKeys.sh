#!/bin/sh

BPATH="./run"
NODES_BASEPATH="${BPATH}/logics"
TENANTID=`cat ${NODES_BASEPATH}/tenantID`
DEVICEID=`cat ${NODES_BASEPATH}/deviceID`

TMPCONFIG="/tmp/genLocalKeys_openssl.cfg"
KEYPATH="${BPATH}/config"
KEYNAME="node"

FQDN="${DEVICEID}.${TENANTID}.msgbus.skidata.net"
ROOTCA="${BPATH}/config/rootCA.pem"
ROOTKEY="${BPATH}/config/rootCA.key"

#-- write openssl configuration
echo "FQDN = ${FQDN}" > ${TMPCONFIG}
echo "ORGNAME = Skidata AG" >> ${TMPCONFIG}
echo "ALTNAMES = DNS:${FQDN}"  >> ${TMPCONFIG}
echo "[ req ]" >> ${TMPCONFIG}
echo "default_bits = 2048" >> ${TMPCONFIG}
echo "default_md = sha256" >> ${TMPCONFIG}
echo "prompt = no" >> ${TMPCONFIG}
echo "encrypt_key = no" >> ${TMPCONFIG}
echo "distinguished_name = dn" >> ${TMPCONFIG}
echo "[ dn ]" >> ${TMPCONFIG}
echo "C = AT" >> ${TMPCONFIG}
echo "L = Grödig" >> ${TMPCONFIG}
echo "O = Skidata AG" >> ${TMPCONFIG}
echo "CN = ${FQDN}" >> ${TMPCONFIG}
echo "[ req_ext ]" >> ${TMPCONFIG}
echo "subjectAltName = " >> ${TMPCONFIG}

openssl genpkey -out ${KEYPATH}/${KEYNAME}.key -outform PEM -algorithm rsa -pkeyopt rsa_keygen_bits:2048
openssl req -new -config ${TMPCONFIG} -key ${KEYPATH}/${KEYNAME}.key -out ${KEYPATH}/${KEYNAME}.csr
openssl x509 -req -in ${KEYPATH}/${KEYNAME}.csr -CA ${ROOTCA} -CAkey ${ROOTKEY} -CAcreateserial -out ${KEYPATH}/${KEYNAME}.crt -days 3650 -sha256
rm ${TMPCONFIG}
rm ${KEYPATH}/${KEYNAME}.csr
