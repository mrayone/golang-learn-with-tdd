package concurrencypatterns

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func fadeIn(ch1, ch2 <-chan string) <-chan string {
	chanIn := make(chan string)
	go func() {
		for {
			chanIn <- <-ch1
		}
	}()

	go func() {
		for {
			chanIn <- <-ch2
		}
	}()
	return chanIn
}

func FanIn() {
	positionChannel := fadeIn(updatePosition("Legolas :"), updatePosition("Gendalf :"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-positionChannel)
	}

	fmt.Println("Done with getting updates on positions.")
}

func FanOut() {
	var myNumbers [10]int
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano()) // gera um array de 10 números randomicos
		myNumbers[i] = rand.Intn(50)
	}

	// converte sequencialmente os valores para string conforme forem chamados
	mychannelOut := channelGenerator(myNumbers)

	// multiplica os valores do canal por 2, conforme estiverem disponíveis
	mychannel1 := double(mychannelOut)
	mychannel2 := double(mychannelOut)

	// junta o output dos dois canais conforme estiverem disponíveis
	// e envia para o canal myChannelIn
	mychannelIn := fanIn(mychannel1, mychannel2)

	// imprime os valores dos canais conforme as operações de sending e receiving
	// forem sendo resolvidas.
	for i := 0; i < len(myNumbers); i++ {
		fmt.Println(<-mychannelIn)
	}
}

func channelGenerator(numbers [10]int) <-chan string {
	channel := make(chan string)
	go func() {
		for _, i := range numbers {
			channel <- strconv.Itoa(i)
		}
		close(channel)
	}()
	return channel
}

func double(inputchannel <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := range inputchannel {
			num, err := strconv.Atoi(i)
			if err != nil {
				// handle error
			}
			channel <- fmt.Sprintf("%d * 2 = %d", num, num*2)
		}
		close(channel)
	}()
	return channel
}

func fanIn(inputchannel1, inputchannel2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case message1 := <-inputchannel1:
				channel <- message1
			case message2 := <-inputchannel2:
				channel <- message2
			}
		}
	}()
	return channel
}
