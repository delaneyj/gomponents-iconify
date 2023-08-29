package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeSlashes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 10l36-6v7L6 17v-7Zm0 14l36-6v7L6 31v-7Zm0 14l36-6v7L6 45v-7Z"/>`),
		g.Group(children),
	)
}