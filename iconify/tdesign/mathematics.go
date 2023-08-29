package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mathematics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.125 1H22v22H2V1.83l10.125 10V1Zm2 20H20v-1.333h-3.15v-2H20v-1.334h-3.15v-2H20V13h-3.15v-2H20V9.667h-3.15v-2H20V6.333h-3.15v-2H20V3h-5.875v18Zm-2 0v-6.36L4 6.615V21h8.125Z"/>`),
		g.Group(children),
	)
}