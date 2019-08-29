create folder fot upload files: mkdir temp-images

run app: go run main.go models.go utils.go

test-curl: curl -X POST -F myfiles=@test.jpg http://localhost:8080
result {"Name":"upload-351526307.png","CropName":"100x100upload-351526307.png"}

test-curl: curl -H 'Content-Type: application/json' -X PUT -d "{\"url\":\"http://i.imgur.com/m1UIjW1.jpg\"}" http://localhost:8080/
result: {"Name":"upload-174286797.png","CropName":"100x100upload-174286797.png"}
