package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingGovernmentTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 2a.5.5 0 0 0-.5.5v2.6A5.002 5.002 0 0 0 5 10v1H4a2 2 0 0 0-2 2v4a1 1 0 0 0 1 1h4.5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v2a.5.5 0 0 0 .5.5H17a1 1 0 0 0 1-1v-4a2 2 0 0 0-2-2h-1v-1a5 5 0 0 0-5-5V4h2.5a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-3ZM6 13.5v2a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 1 0Zm8.5-.5a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 .5-.5ZM9 9.5v2a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 1 0Zm2.5-.5a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}