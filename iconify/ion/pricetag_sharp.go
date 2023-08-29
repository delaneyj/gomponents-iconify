package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PricetagSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M304 32L16 320l176 176l288-288V32Zm80 128a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}