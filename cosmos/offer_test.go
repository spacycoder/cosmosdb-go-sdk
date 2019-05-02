package cosmos

import "testing"

func TestOffer(t *testing.T) {
	client := getDummyClient()
	offer := client.Offer("myOffer")

	if offer.client.rType != "offers" {
		t.Errorf("Wrong resource type %s", offer.client.rType)
	}

	if offer.client.rLink != "myOffer" {
		t.Errorf("Wrong resource link: %s should be: %s", offer.client.rLink, "myOffer")
	}

	if offer.client.path != "offers/myOffer" {
		t.Errorf("Wrong path %s", offer.client.path)
	}

	offers := client.Offers()
	if offers.client.rType != "offers" {
		t.Errorf("Wrong resource type: %s should be: %s", offers.client.rType, "offers")
	}

	if offers.client.rLink != "" {
		t.Errorf("Wrong resource link: %s should be: %s", offers.client.rLink, "")
	}

	if offers.client.path != "offers" {
		t.Errorf("Wrong path %s", offers.client.path)
	}
}
