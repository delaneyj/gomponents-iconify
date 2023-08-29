package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M182 184a30 30 0 0 1-30 30c-12.9 0-25.36-8.38-29.63-19.92a6 6 0 0 1 11.26-4.16C136.13 196.69 144.2 202 152 202a18 18 0 0 0 0-36H40a6 6 0 0 1 0-12h112a30 30 0 0 1 30 30ZM150 72a30 30 0 0 0-30-30c-12.9 0-25.36 8.38-29.63 19.92a6 6 0 1 0 11.26 4.16C104.13 59.31 112.2 54 120 54a18 18 0 0 1 0 36H24a6 6 0 0 0 0 12h96a30 30 0 0 0 30-30Zm58 2c-12.9 0-25.36 8.38-29.63 19.92a6 6 0 1 0 11.26 4.16C192.13 91.31 200.2 86 208 86a18 18 0 0 1 0 36H32a6 6 0 0 0 0 12h176a30 30 0 0 0 0-60Z"/>`),
		g.Group(children),
	)
}