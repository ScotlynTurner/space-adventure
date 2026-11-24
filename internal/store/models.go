package store

type User struct {
	UserName string `firestore:"userName"`
}

type Achievement struct {
	Name string `firestore:"name"`
	Icon string `firestore:"icon"`
}

type Choice struct {
	Label      string `firestore:"label"`
	NextNodeID string `firestore:"nextNodeId"`
}

// StoryNode fields are the shared shape for story documents.
// The Firestore document ID identifies the node; it is not duplicated here.
type StoryNode struct {
	Description string   `firestore:"description"`
	Image       string   `firestore:"image"`
	Choices     []Choice `firestore:"choices"`
	EndingID    string   `firestore:"endingId,omitempty"`
}
