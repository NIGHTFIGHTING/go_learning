package client

import (
    "ch15/series"
    "testing"
)


func TestPackaget(t *testing.T) {
    t.Log(series.GetbonacciSeries(5))
    t.Log(series.Square(5))
}
