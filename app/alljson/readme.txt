STRUCTURE
$ migrate -source file:///home/henry/work/src/gotodo/app/migrations -database postgres://sysdba:orbdba@localhost:5432/tododb down

$ migrate -source file:///home/henry/work/src/gotodo/app/migrations -database postgres://sysdba:orbdba@localhost:5432/tododb up

(Force)
$ migrate -source file:///home/henry/work/src/gotodo/app/migrations -database postgres://sysdba:orbdba@localhost:5432/tododb force 1519742319