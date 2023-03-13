package main

import (
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatUtilCanConvertAListOfOrganisationsToJson(t *testing.T) {
	org1 := safetorun.ListOrganisationsListOrganisationsOrganisationListItemsOrganisation{
		OrganisationId: "abc",
	}

	org2 := safetorun.ListOrganisationsListOrganisationsOrganisationListItemsOrganisation{
		OrganisationId: "def",
	}

	organisationsList := []safetorun.ListOrganisationsListOrganisationsOrganisationListItemsOrganisation{org1, org2}

	expectedJson := "[{\"OrganisationId\":\"abc\"},{\"OrganisationId\":\"def\"}]"

	actualJson := ToJson(organisationsList)

	assert.Equal(t, expectedJson, actualJson)
}
