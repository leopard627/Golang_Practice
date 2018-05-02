package channels

import (
	"fmt"
	"testing"
)

func TestBufferingChannels(t *testing.T) {

	/*
		채널은 _버퍼링_될 수 있습니다. make 에 두번째 인자로 버퍼 용량을 넣음으로써 해당 용량만큼 버퍼링되는 채널을 생성할 수 있습니다:

		ch := make(chan int, 100)
		버퍼링되는 채널로의 송신은 버퍼가 꽉 찰 때까지 블록됩니다. 수신측은 버퍼가 빌 때 블록됩니다.

		예제를 수정해서 버퍼를 넘치게 해보고 어떻게 동작하는지 확인해 보세요.

	*/

	c := make(chan int, 2)
	c <- 1
	c <- 20000000000000000000

	fmt.Println(<-c)
	fmt.Println(<-c)

}
