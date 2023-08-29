package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinusLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m204.24 60.24l-144 144a6 6 0 0 1-8.48-8.48l144-144a6 6 0 0 1 8.48 8.48ZM66 112a6 6 0 0 0 12 0V78h34a6 6 0 0 0 0-12H78V32a6 6 0 0 0-12 0v34H32a6 6 0 0 0 0 12h34Zm158 66h-80a6 6 0 0 0 0 12h80a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}