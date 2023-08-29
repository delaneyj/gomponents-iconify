package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 32a12 12 0 0 1-12 12H56a12 12 0 0 1 0-24h144a12 12 0 0 1 12 12Zm-32 48v144a20 20 0 0 1-20 20H96a20 20 0 0 1-20-20V80a20 20 0 0 1 20-20h64a20 20 0 0 1 20 20Zm-24 4h-56v136h56Z"/>`),
		g.Group(children),
	)
}