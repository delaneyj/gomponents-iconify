package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Zm0 14a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Zm1-8a1 1 0 1 0 0 2h20a1 1 0 1 0 0-2H4Z"/>`),
		g.Group(children),
	)
}