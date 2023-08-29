package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1a5 5 0 1 0 0 10A5 5 0 0 0 6 1ZM2 6a4 4 0 0 1 6.453-3.16L2.84 8.453A3.983 3.983 0 0 1 2 6Zm1.547 3.16L9.16 3.547A4 4 0 0 1 3.547 9.16Z"/>`),
		g.Group(children),
	)
}