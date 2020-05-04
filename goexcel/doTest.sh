#!/usr/bin/env bats

@test "セル値出力" {
    run go run main.go get Book1.xlsx Sheet1 A1
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1000" ]
}

@test "シートダンプ" {
    run go run main.go get Book1.xlsx Sheet1
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1000," ]
}

@test "シート一覧" {
    run go run main.go get Book1.xlsx
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "Sheet1" ]
    [ "${lines[1]}" = "Sheet2" ]
}

@test "ヘルプ" {
    run go run main.go get -h
    [ "$status" -eq 0 ]
}

@test "リソース指定なし(get)" {
    run go run main.go get
    [ "$status" -eq 1 ]
}

@test "存在しないExcelファイルを指定(get)" {
    run go run main.go get a
    [ "$status" -eq 1 ]
}

@test "存在しないシートを指定(get)" {
    run go run main.go get Book1.xlsx a
    [ "$status" -eq 1 ]
}

@test "存在しないセルを指定(get)" {
    run go run main.go get Book1.xlsx Sheet1 a
    [ "$status" -eq 1 ]
}

@test "リソース指定なし(set)" {
    run go run main.go set
    [ "$status" -eq 1 ]
}

@test "存在しないExcelファイルを指定(set)" {
    run go run main.go set a
    [ "$status" -eq 1 ]
}

@test "存在しないシートを指定(set)" {
    run go run main.go set Book1.xlsx a b
    [ "$status" -eq 1 ]
}

@test "存在しないセルを指定(set)" {
    run go run main.go set Book1.xlsx Sheet1 a 1000
    [ "$status" -eq 1 ]
}

@test "シート名更新" {
    run go run main.go set Book1.xlsx Sheet1 Sheet99
    [ "$status" -eq 0 ]
}

@test "更新後シート一覧" {
    run go run main.go get Book1.xlsx
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "Sheet99" ]
    [ "${lines[1]}" = "Sheet2" ]
}

@test "シート名戻し" {
    run go run main.go set Book1.xlsx Sheet99 Sheet1
    [ "$status" -eq 0 ]
}

@test "戻し後シート一覧" {
    run go run main.go get Book1.xlsx
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "Sheet1" ]
    [ "${lines[1]}" = "Sheet2" ]
}

@test "セル値更新" {
    run go run main.go set Book1.xlsx Sheet1 A1 10000
    [ "$status" -eq 0 ]
}

@test "更新後セル値出力" {
    run go run main.go get Book1.xlsx Sheet1 A1
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "10000" ]
}

@test "セル値戻し" {
    run go run main.go set Book1.xlsx Sheet1 A1 1000
    [ "$status" -eq 0 ]
}

@test "戻し後セル値出力" {
    run go run main.go get Book1.xlsx Sheet1 A1
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1000" ]
}
