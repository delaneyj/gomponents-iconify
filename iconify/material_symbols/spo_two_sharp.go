package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpoTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 20v-6h4.5v6H11Zm1.5-1.5H14v-3h-1.5v3ZM17 22v-3.75h3v-.75h-3V16h4.5v3.75h-3v.75h3V22H17Zm-8-.05q-3.075-.35-5.038-2.638T2 13.8q0-2.5 1.988-5.438T10 2q3.3 2.8 5.238 5.3T17.75 12H9v9.95Z"/>`),
		g.Group(children),
	)
}