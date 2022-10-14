package safetorun

import "context"

func (client Client) RetrieveLastEventForLinkId(linkId string) (*GetEventsForLinkIdEventsForLinkIdEventsItemsEvent, error) {
	ctx := context.Background()

	response, err := GetEventsForLinkId(ctx, client.GqlClient, linkId)

	if err != nil {
		return nil, err
	}

	event := response.EventsForLinkId.GetItems()[0]
	return &event, nil
}
