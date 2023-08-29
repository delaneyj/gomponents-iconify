package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M442.2 93.1L256 279.3L69.8 93.1L0 162.9l256 256l256-256z"/>`),
		g.Group(children),
	)
}