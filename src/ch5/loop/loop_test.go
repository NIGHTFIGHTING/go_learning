// 08条件和循环

package loop_test

import "testing"

func TestWhiltLoop(t *testing.T) {
    n := 0
    for n < 5 {
        t.Log(n);
        n++;
    }
}
