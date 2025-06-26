# Go with DDD

## Create Schema
```shell
go run -mod=mod entgo.io/ent/cmd/ent new User
```

## Generate Codes
```shell
```

## Generate Mock

```shell
# mockgen -source=[path] -destination=[path] -package=[packageName]
mockgen -source=application/required/emailSender.go -destination=application/required/mockEmailSender.go -package required
```