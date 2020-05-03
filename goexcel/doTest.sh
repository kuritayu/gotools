#!/usr/bin/env bats

@test "セル値出力" {
    run go run main.go get -c Book1.xlsx Sheet1 A1
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1000" ]
}

@test "シートダンプ" {
    run go run main.go get -d Book1.xlsx Sheet1 ","
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1000," ]
}

@test "シート一覧" {
    run go run main.go get -s Book1.xlsx
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "Sheet1" ]
    [ "${lines[1]}" = "Sheet2" ]
}

@test "ヘルプ" {
    run go run main.go get -h
    [ "$status" -eq 0 ]
}

@test "リソース指定なし" {
    run go run main.go get
    [ "$status" -eq 1 ]
}