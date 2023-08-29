package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2.5a.5.5 0 0 0-1 0v15a.5.5 0 0 0 1 0v-15ZM2 6a2 2 0 0 1 2-2h4v12H4a2 2 0 0 1-2-2V6Zm9 10h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4v12Z"/>`),
		g.Group(children),
	)
}