package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedPublicSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M.5 8a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm4.5.25v3a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1H7.5v-1.5a1.5 1.5 0 0 1 2.443-1.167a.749.749 0 1 0 .943-1.166A3 3 0 0 0 6 5.75v1.5a1 1 0 0 0-1 1Z"/>`),
		g.Group(children),
	)
}