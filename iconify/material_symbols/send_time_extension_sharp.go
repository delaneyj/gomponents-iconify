package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendTimeExtensionSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22v-4l4-1l-4-1v-4l10 5l-10 5Zm-4.2-1H3v-5.8q1.2 0 2.1-.762T6 12.5q0-1.175-.9-1.937T3 9.8V4h6q0-1.05.725-1.775T11.5 1.5q1.05 0 1.775.725T14 4h6v9.25l-9-4.5v9.3q-1 .2-1.6.938T8.8 21Z"/>`),
		g.Group(children),
	)
}