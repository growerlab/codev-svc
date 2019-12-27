## git service

### golang
depends on golang greater than 1.13

### dependencies
- libgit2/git2go & libgit2/libgit2
```
git clone http://github.com/libgit2/libgit2 -b v0.27.9 --depth=1
cd libgit2 && mkdir build && cd build
cmake build .. && make install
pkg-config --libs --cflags libgit2 # check if installed successfully or `pkg-config --modversion libgit2`
```
- gin-gonic/gin
- graphql-go/graphql
- joho/godotenv

### installation
``` bash
go get github.com/growerlab/codev-svc
cd $GOPATH/src/github.com/growerlab/codev-svc
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