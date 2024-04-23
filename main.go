package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const cacheSize = 10

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Lenght int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{
		Head: head,
		Tail: tail,
	}
}

func (cache *Cache) Check(str string) {
	node := &Node{}

	if val, ok := cache.Hash[str]; ok {
		node = cache.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	cache.Add(node)
	cache.Hash[str] = node
}

func (cache *Cache) Remove(node *Node) *Node {
	log.Warn().Msgf("удаление: %v", node.Val)
	left := node.Left
	right := node.Right

	right.Left = left
	left.Right = right

	cache.Queue.Lenght--
	delete(cache.Hash, node.Val)
	return node
}

func (cache *Cache) Add(node *Node) {
	log.Info().Msgf("добавление: %v", node.Val)
	tmp := cache.Queue.Head.Right

	cache.Queue.Head.Right = node
	node.Left = cache.Queue.Head
	node.Right = tmp
	tmp.Left = node

	cache.Queue.Lenght++
	if cache.Queue.Lenght > cacheSize {
		cache.Remove(cache.Queue.Tail.Left)
	}
}

func (cache *Cache) Display() {
	cache.Queue.Display()
}

func (queue *Queue) Display() {
	node := queue.Head.Right
	fmt.Printf("%v - [", queue.Lenght)
	for i := 0; i < queue.Lenght; i++ {
		fmt.Printf("{%v}", node.Val)
		if i < queue.Lenght-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Printf("]\n")
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Заупск кеширования")
	cache := NewCache()
	for _, word := range []string{
		"лев",
		"тигр",
		"слон",
		"жираф",
		"лев",
		"зебра",
		"кенгуру",
		"пингвин",
		"лев",
		"коала",
		"панда",
		"белый медведь",
		"горилла",
		"лев",
		"бегемот",
		"носорог",
		"крокодил",
		"лев",
		"аллигатор",
		"акула",
		"осьминог",
		"черепаха",
		"лев",
		"змея",
		"паук",
	} {
		cache.Check(word)
		cache.Display()
	}
}
