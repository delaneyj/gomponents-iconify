package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdentityPlatform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13q1.45 0 2.475-1.025Q15.5 10.95 15.5 9.5q0-1.45-1.025-2.475Q13.45 6 12 6q-1.45 0-2.475 1.025Q8.5 8.05 8.5 9.5q0 1.45 1.025 2.475Q10.55 13 12 13Zm0 9.5L3 17V7l9-5.5L21 7v10Zm0-2.35l5.6-3.425q-1.2-.825-2.612-1.275Q13.575 15 12 15t-2.987.45Q7.6 15.9 6.4 16.725Z"/>`),
		g.Group(children),
	)
}