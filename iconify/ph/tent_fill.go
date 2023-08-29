package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m255.31 196.75l-64-144A8 8 0 0 0 184 48H72a8 8 0 0 0-7.31 4.75v.12l-64 143.88A8 8 0 0 0 8 208h240a8 8 0 0 0 7.31-11.25ZM64 192H20.31L64 93.7Zm16 0V93.7l43.69 98.3Z"/>`),
		g.Group(children),
	)
}