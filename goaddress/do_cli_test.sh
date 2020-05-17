#!/usr/bin/env bats

# 正常
@test "update" {
    run goaddress update
    [ "$status" -eq 0 ]
}

@test "search" {
    run goaddress search 1810002
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1810002 東京都三鷹市牟礼" ]
}

@test "search2" {
    run goaddress search 181-0002
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1810002 東京都三鷹市牟礼" ]
}

@test "search_reverse" {
    run goaddress search -r 東京都三鷹市牟礼
    [ "$status" -eq 0 ]
    [ "${lines[0]}" = "1810002 東京都三鷹市牟礼" ]
}

@test "show help1" {
    run goaddress -h
    [ "$status" -eq 0 ]
}

@test "show help2" {
    run goaddress update -h
    [ "$status" -eq 0 ]
}

@test "show help3" {
    run goaddress search -h
    [ "$status" -eq 0 ]
}

@test "show help4" {
    run goaddress search -h -r
    [ "$status" -eq 1 ]
}
