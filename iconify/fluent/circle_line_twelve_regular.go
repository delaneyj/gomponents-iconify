package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLineTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1a5 5 0 1 0 0 10A5 5 0 0 0 6 1ZM2 6a4 4 0 1 1 8 0H2Zm.126 1h7.748a4.002 4.002 0 0 1-7.748 0Z"/>`),
		g.Group(children),
	)
}