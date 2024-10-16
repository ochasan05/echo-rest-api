# デプロイ方法（ECR）
- docker build --platform linux/x86_64 -t stg-go-app-01 .
- docker tag stg-go-app-01:latest 303944428269.dkr.ecr.ap-northeast-1.amazonaws.com/stg-go-app-01:latest
- docker push 303944428269.dkr.ecr.ap-northeast-1.amazonaws.com/stg-go-app-01:latest