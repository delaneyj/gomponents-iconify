package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpoTwoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 20v-6h4.5v6H11Zm1.5-1.5H14v-3h-1.5v3ZM17 22v-3.75h3v-.75h-3V16h4.5v3.75h-3v.75h3V22H17Zm-8-.05q-3.075-.35-5.038-2.638T2 13.8q0-2.5 1.988-5.438T10 2q3.3 2.8 5.238 5.3T17.75 12h-2.075q-.55-1.575-1.975-3.425T10 4.65Q7.025 7.375 5.513 9.675T4 13.8q0 2.4 1.388 4.088T9 19.925v2.025Zm.85-8.15Z"/>`),
		g.Group(children),
	)
}