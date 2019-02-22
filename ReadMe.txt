A simple demo which demonstates a simple chat using libMsgbus from Go

Generate keyfiles:
1) mkdir -p run/{config|logics}
2) echo "mydeviceID" > logics/deviceid
3) echo "mytenantID" > logics/tenantid
4) ./genLocalKeys.sh

Build project:
- go build

Run:
- ./GoMsgBusDemo <blname>
