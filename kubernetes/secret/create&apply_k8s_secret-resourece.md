#### 1 create secret and apply it to pod
##### create secret  by kind:Secret
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: secret-privatekey
type: Opaque
data:
private_key.pem: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBd1Exc2pmelBkTUw2Znh6UklnVmZoYTZON0tialdTK2NOcUV3QmNTTkYxWXZhMmhKCjZ6Z1hqRkFzZHE5bHBWRDFDL1VZNTNYNWJHM0N3MnowdWdYbVlHZllkNnVJd1dvMEh2QVZuMFA2T2NVTnR1SEwKMzVzSTdPRGNpK1J5Z3VSRjBDNWZQTXhjcW42djB0azk0dTd4L3FiZnQrWExpb21mL1ZXV0pQVFlNM2toRGh3MworcTA4UUlxdEZTRjh0bEtDUElyQjBieVM0Rm8zeDdsZEUxSkg1ZFM5SHNCMFA5a0xMWEROYWNrTUh0bkNMaUZwClQ3MXdsU0c1Mm1vNFBsQ2NkR0xuaDFvd01zNTVwYkNwa1ZTZUE3bDNQUnlBdWduZXdMbGR2bjdUNVlFNjZZTzMKOVg5SGNxNW8zN3hLa05ud0hDMm5RK0Z0V2xnYjZrem5teURNOXdJREFRQUJBb0lCQVFDdjBHN3RmTERla0hlYgpiZjRVTXJxRXY2eW5PbkhRbG1oNDVDRWRENXpEQlEyWWp0aks3RUdkMnFJejBKY01rSWNxeGFOUm9JSndPaUhtCjRvS2FLNmRjWXhha0hjY2xCbmpET0RrbzI3cTJBL3p5Y2Y2Ky9LMkxOVm9GMlI1a2tFbjRSMU1heHE4WjR0aHoKZUw5QXZnWUx1YVFERWJkbHl1SzJ6OHcyaWZtZ0hYZVY2V2tvV050cHU5Ni8vTXlpNnpYQ1YzKzEwUnNIZWc3Ngo2VElBU0NQVXBCbk1QZnI2bGJYQktod2gycTBSTWhMSGduMVA5aTAvM3RIYUNGMmdPcm9FZzNVL1lkU2ZSc0dFCnpOamtDZFR2OEhnTU1jTkVuc1hGdzRVYm9PcFJodlBmMHNub2tITHlKS0J0QWF0eUhZYmJORTNuVlRON3ZyZ2YKK2hDN2d6aXBBb0dCQVBybndlUnpjbndBNllOQWluQlFIczBmbGRWV2dhYlU2U2lRRTVVNmx0MmhnbWMybzlRegpsQUNDL0U1RFFmY3RMQ0V2LzVtQndPcko3Q3ZnT2dDUVIrWkw5QUdRcWlyRjFBNkhUVUgwY3BhM1pjMWV5WHJVCm9PQVltMUJ2T2tyZXlWcHE2SHhjMkswRkNGLzg4aUZqajV1VkVFVVVOSUtDQXRoMDZpbkNnZlBWQW9HQkFNVDQKOEYvaDAxcEZTU2ZrYjZ2Sld4RUFrdmVWcjhnVHp6amxRR2Y1QXd0RVJDTHc3VkszUGdLZXI2WEdWdGQ0MVlvegpwc0JpRVlSczRxdnN6aDc2Nm9rOGZNVkNYOHZYeWtjS1BIT2hzVG5zcEZXbjhtUFhLM3gweVpjRWY2Q3BlNndXCkFlbXBxekwxTUV0TkhBNTlrMCtQNDhlQnhSODgyRTZubE1zWHpmK2JBb0dBSWdRTG1HSkNjaWRaZ1M3ZDFlNDIKenM1cWJOcm1odXkwazRnODcyMWJDTDhkdzhwM21ZeElrMjB1c1ptU3R1VUw1NC92VWl0eU1TS2cveTNPRDBlQQpSK3VpTUJnaHlkeTZMQ0lSeWxCT3ZMb3VkaEpVdEc5aFJDQng1Z3krVldvdzJDNUlTSnY1MERNdmVIdjlnNk5RCjArSDRxN1RhalpyOHNjWGYvVHRlak9VQ2dZQXlhaHpzRFpyUThnYmxaUHlJRllOdmVKd2xMblROV3ZTZzlWeGsKd0VGZE10M1ZxNkN5bVNBUC81bXBibmh2c2dmRjFhNktjdzlVdTZIUXEwMmVkRTV2VGNJSm94RnQxUTk2MjAzWgpzcnJ2dm5mWlRLRW5tTDBTbjdteEkzK2ZHWUlENjZZVnJrMlpQMVJiRWFOcXVnMW9RY1hsSEh2ZG9POGtRcFE4CjN4emtuUUtCZ1FEUlV4K3c0T3Y2ZlVVMkk1OW1MQlR6eXVSL2tpSXVYY29JTHVFVC9QVUdqaGlVd0JVY0YrVloKTnFDZytwQ3l6Z0Q1ZGNjZE93dVRRdEJVUnRIVnhBNFVXRVduRi92NklFSi9nVm9VS25seG1oS3BCUjFUV1FNawpSUVg2d0gwU0JUUzVvYVBkN1dTSmNpUm01ZXZSRkNxaVlhZTlYdWpZbHp2ZkozUlV1dlNyL1E9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo

```
##### apply secret by volume
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-secret-pod
spec:
  containers:
    - name: test-secret-pod
      image: alpine
      command: ["/bin/sh","-c","ping 8.8.8.8"]
      volumeMounts:
        - name: secret-volume
          mountPath: /etc/secret
          readOnly: true
  volumes:
    - name: secret-volume
      secret:
        secretName: secret-privatekey
```
#### 2 refer to:
1 [secret](https://blog.csdn.net/Victor2code/article/details/106042691)



 
