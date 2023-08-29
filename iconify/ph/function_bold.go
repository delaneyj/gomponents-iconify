package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 40a12 12 0 0 1-12 12h-29.29A20 20 0 0 0 151 68.42L142.38 116H184a12 12 0 0 1 0 24h-46l-9.44 51.87A44 44 0 0 1 85.29 228H56a12 12 0 0 1 0-24h29.29A20 20 0 0 0 105 187.58l8.62-47.58H72a12 12 0 0 1 0-24h46l9.44-51.87A44 44 0 0 1 170.71 28H200a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}