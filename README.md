## git service

### installation
``` bash
go get github.com/growerlab/svc
cd $GOPATH/src/github.com/growerlab/svc
go mod download
```

### build and run
`APP_ENV=dev go run main.go`

### env
- dev
- staging
- test
- production

### graphql
http://your-server-ip:port/graphql


### TODO
- [x] Create repo
- [ ] Delete repo
- [ ] Fork repo
- [x] Create branch
- [x] Branch list
- [ ] Default branch
- [x] Delete branch
- [ ] Rename branch
- [ ] Diff branch
- [ ] File list
- [ ] Delete file
- [ ] Tag list
- [ ] Delete Tag
- [ ] Commit list
- [ ] Create commit (create file)
- [ ] Diff commit
- [ ] Commit desc
- [ ] Push Hook
- [ ] README
