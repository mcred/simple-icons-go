package simple_icons_go

import (
	"testing"
)

func TestIcons_getByNameFails(t *testing.T) {
	icon := testIcons.getByName("invalid_slug")
	if icon.Title != "" {
		t.Errorf("Icon.Get() got = %v, expected %v", icon.Title, "")
	}
}
