package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1a5 5 0 1 1 0 10A5 5 0 0 1 6 1Zm-.5 2.5A.5.5 0 0 0 5 4v2.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 0-1H6V4a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}