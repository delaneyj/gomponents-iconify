package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowGraphUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M192 128l-61.8 61.8 93.5 98.2 107-106.7L480 384 330.7 256l-107 112L92.8 227.2 32 288V128z" fill="currentColor"/>`),
		g.Group(children),
	)
}