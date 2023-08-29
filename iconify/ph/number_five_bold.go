package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFiveBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M180 160a60 60 0 0 1-100 44.72a12 12 0 1 1 16-17.88a36 36 0 1 0 .69-54.28a12 12 0 0 1-19.54-11.49l15.08-75.42A12 12 0 0 1 104 36h64a12 12 0 0 1 0 24h-54.16l-8.36 41.79A60 60 0 0 1 180 160Z"/>`),
		g.Group(children),
	)
}