package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m251.66 198.38l-64-144A4 4 0 0 0 184 52H72a4 4 0 0 0-3.63 2.35v.06l-64 143.93A4 4 0 0 0 8 204h240a4 4 0 0 0 3.66-5.62ZM68 74.85V196H14.16ZM76 196V74.85L129.84 196Zm62.6 0L78.16 60H181.4l60.44 136Z"/>`),
		g.Group(children),
	)
}