# gorm with mysql

## tl;dr

- [ ] Associations
  - has one, has many, belongs to, many to many, polymorphism, single-table inheritance
- [ ] Hooks
  - before/after create/save/update/delete/find
- [ ] Eager loading with Preload, Joins
- [ ] Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
- [ ] Prometheus


## install
```console
go get -u github.com/k0kubun/pp/v3 # pretty print
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```