package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.646 5.902a.5.5 0 0 0 .708 0L5.5 4.756V8.5a.5.5 0 0 0 1 0V4.756l1.146 1.146a.5.5 0 1 0 .708-.707l-2-2a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0 0 .707ZM6 1a5 5 0 1 0 0 10A5 5 0 0 0 6 1ZM2 6a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/>`),
		g.Group(children),
	)
}