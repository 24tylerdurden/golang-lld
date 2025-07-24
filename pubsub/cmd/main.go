package main

import (
	"pubsub/pubsub"
	"time"
)

func main() {
	ps := pubsub.NewPubSub()

	// topics
	politics := pubsub.NewTopic(1, "politics")
	sports := pubsub.NewTopic(2, "sports")

	// subscribers
	alice := pubsub.NewPoliticalNewsSubscriber("alice")
	bob := pubsub.NewSportsNewsSubscriber("bob")

	alice.Listen()
	bob.Listen()
	ps.AddSubscriberToTopic(politics, alice)
	ps.AddSubscriberToTopic(sports, bob)

	// publishers
	poPublisher := &pubsub.PoliticalNewsPublisher{Pubsub: ps}
	spPublisher := &pubsub.SportsNewsPublisher{Pubsub: ps}

	poPublisher.PublishToTopic(politics, pubsub.NewEvent("Election Results are announced"))
	spPublisher.PublishToTopic(sports, pubsub.NewEvent("Team A won the match"))

	time.Sleep(2 * time.Second)

}
