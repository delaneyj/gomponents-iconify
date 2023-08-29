package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewQuiltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5v13h17V5H4m2 11V7h3v9H6m5 0v-3.5h3V16h-3m8 0h-3v-3.5h3V16m-8-5.5V7h8v3.5h-8Z"/>`),
		g.Group(children),
	)
}