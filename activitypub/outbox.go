package activitypub

type (
	// OutboxStream contains activities the user has published,
	// subject to the ability of the requestor to retrieve the activity (that is,
	// the contents of the outbox are filtered by the permissions of the person reading it).
	OutboxStream Outbox

	// Outbox is a type alias for an Ordered Collection
	Outbox OrderedCollection
)