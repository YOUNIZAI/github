### 开发平台(linux+git+gitlab)
 ```
 - git
   -- code version
   -- .gitignore
   -- git config
 -gitlab
   -- .gitlab-ci.yaml
   -- gitlab: merge->comment->e-mail->jira;build images;save images;auto test 
 ```
### 版本发布(Dockerfile+docker+gitlab images)
```
1.docker build via dockerfile
2.docker push to gitlab images hub
```
### 测试平台(tester,ci-cd,k8s)
```
1.tester provide test case
2.ci-cd auto run test case
3.k8s provide test environment
```
