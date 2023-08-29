package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractCircleTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 6a5 5 0 1 1 10 0A5 5 0 0 1 1 6Zm3-.5a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1H4Z"/>`),
		g.Group(children),
	)
}