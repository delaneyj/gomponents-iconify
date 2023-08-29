package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSplitTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.003 2.381a1.5 1.5 0 0 0-2.006 0l-5.5 4.95A1.5 1.5 0 0 0 3 8.446V15.5A1.5 1.5 0 0 0 4.5 17h11a1.5 1.5 0 0 0 1.5-1.5V8.446a1.5 1.5 0 0 0-.497-1.115l-5.5-4.95Zm-.503 11.12v1a.5.5 0 0 1-1 0v-1a.5.5 0 1 1 1 0ZM10 9a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-1 0v-1A.5.5 0 0 1 10 9Zm.5-3.5v1a.5.5 0 0 1-1 0v-1a.5.5 0 0 1 1 0Z"/>`),
		g.Group(children),
	)
}