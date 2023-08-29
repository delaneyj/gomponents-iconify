package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingGovernmentThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.667 2H16a1 1 0 0 0-1 1v4.055A9 9 0 0 0 7 16H5a3 3 0 0 0-3 3v9a1 1 0 0 0 1 1h10v-4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4h10a1 1 0 0 0 1-1v-9a3 3 0 0 0-3-3h-2a9 9 0 0 0-8-8.945V6h3.667a.667.667 0 0 0 .667-.667V2.667A.667.667 0 0 0 20.667 2ZM12 13a1 1 0 0 1 1 1v3a1 1 0 1 1-2 0v-3a1 1 0 0 1 1-1Zm8 0a1 1 0 0 1 1 1v3a1 1 0 1 1-2 0v-3a1 1 0 0 1 1-1Zm-3 1v3a1 1 0 1 1-2 0v-3a1 1 0 1 1 2 0Zm-9 6a1 1 0 0 1 1 1v3a1 1 0 0 1-2 0v-3a1 1 0 0 1 1-1Zm16 0a1 1 0 0 1 1 1v3a1 1 0 0 1-2 0v-3a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}