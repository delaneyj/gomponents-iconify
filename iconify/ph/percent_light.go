package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m204.24 60.23l-144 144a6 6 0 0 1-8.48-8.48l144-144a6 6 0 1 1 8.48 8.49ZM52 100a34 34 0 1 1 24 10a33.78 33.78 0 0 1-24-10Zm2-24a22 22 0 1 0 6.44-15.56A21.86 21.86 0 0 0 54 76Zm160 104a34 34 0 1 1-10-24a33.78 33.78 0 0 1 10 24Zm-12 0a21.87 21.87 0 0 0-6.44-15.56A22 22 0 1 0 202 180Z"/>`),
		g.Group(children),
	)
}