package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BroadActivityFeedTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 3A1.5 1.5 0 0 0 2 4.5v4A1.5 1.5 0 0 0 3.5 10h13A1.5 1.5 0 0 0 18 8.5v-4A1.5 1.5 0 0 0 16.5 3h-13Zm0 9A1.5 1.5 0 0 0 2 13.5v2A1.5 1.5 0 0 0 3.5 17h4A1.5 1.5 0 0 0 9 15.5v-2A1.5 1.5 0 0 0 7.5 12h-4Zm9 0a1.5 1.5 0 0 0-1.5 1.5v2a1.5 1.5 0 0 0 1.5 1.5h4a1.5 1.5 0 0 0 1.5-1.5v-2a1.5 1.5 0 0 0-1.5-1.5h-4Z"/>`),
		g.Group(children),
	)
}