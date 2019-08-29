<strong>create folder fot upload files:</strong> mkdir temp-images

<strong>run app:</strong> go run main.go models.go utils.go

<strong>test-curl:</strong> curl -X POST -F myfiles=@test.jpg http://localhost:8080
<strong>result:</strong> {"Name":"upload-351526307.png","CropName":"100x100upload-351526307.png"}

<strong>test-curl:</strong> curl -H 'Content-Type: application/json' -X PUT -d "{\"url\":\"http://i.imgur.com/m1UIjW1.jpg\"}" http://localhost:8080/
<strong>result:</strong> {"Name":"upload-174286797.png","CropName":"100x100upload-174286797.png"}
