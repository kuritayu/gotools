#!/usr/bin/env bats

# build
setup() {
    go build -o goexcel ./main.go
    mv goexcel ${GOPATH}/bin
}

# 正常
@test "create book" {
    run goexcel create Book1.xlsx
    [ "$status" -eq 0 ]
}

@test "create sheet" {
    run goexcel create Book1.xlsx Sheet2
    [ "$status" -eq 0 ]
}

@test "create sheet2" {
    run goexcel create Book1.xlsx Sheet3
    [ "$status" -eq 0 ]
}

@test "update sheet name" {
    run goexcel set Book1.xlsx Sheet3 Sheet99
    [ "$status" -eq 0 ]
}

@test "set cell" {
    run goexcel set Book1.xlsx Sheet1 A1 1000
    [ "$status" -eq 0 ]
}

@test "get cell" {
    run goexcel get Book1.xlsx Sheet1 A1
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1000" ]
}

@test "dump sheet" {
    run goexcel get Book1.xlsx Sheet1
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1000," ]
}

@test "get sheet list" {
    run goexcel get Book1.xlsx
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "Sheet1" ]
    [ "${lines[1]}" = "Sheet2" ]
}

@test "delete sheet" {
    run goexcel delete Book1.xlsx Sheet2
    [ "$status" -eq 0 ]
}

@test "delete book" {
    run goexcel delete Book1.xlsx
    [ "$status" -eq 0 ]
}

@test "show help" {
    run goexcel get -h
    [ "$status" -eq 0 ]
}

# 異常
@test "no argument for get subcommand" {
    run goexcel get
    [ "$status" -eq 1 ]
}

@test "no argument for set subcommand" {
    run goexcel set
    [ "$status" -eq 1 ]
}

@test "no argument for create subcommand" {
    run goexcel create
    [ "$status" -eq 1 ]
}

@test "no argument for delete subcommand" {
    run goexcel delete
    [ "$status" -eq 1 ]
}

@test "noonexistent file for get subcommand" {
    run goexcel get a
    [ "$status" -eq 1 ]
}

@test "noonexistent file for set subcommand" {
    run goexcel set a
    [ "$status" -eq 1 ]
}

@test "noonexistent file for create subcommand" {
    run goexcel set a b
    [ "$status" -eq 1 ]
}

@test "noonexistent file for delete subcommand" {
    run goexcel set a
    [ "$status" -eq 1 ]
}

@test "noonexistent sheet for get subcommand" {
    run goexcel get Book1.xlsx a
    [ "$status" -eq 1 ]
}

@test "noonexistent sheet for set subcommand" {
    run goexcel set Book1.xlsx a b
    [ "$status" -eq 1 ]
}

@test "noonexistent sheet for delete subcommand" {
    run goexcel set Book1.xlsx a
    [ "$status" -eq 1 ]
}

@test "noonexistent cell for get subcommand" {
    run goexcel get Book1.xlsx Sheet1 a
    [ "$status" -eq 1 ]
}

@test "noonexistent cell for set subcommand" {
    run goexcel set Book1.xlsx Sheet1 a 1000
    [ "$status" -eq 1 ]
}
