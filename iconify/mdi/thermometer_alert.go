package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 13V7h2v6h-2m0 4v-2h2v2h-2m-4-4V5c0-1.7-1.3-3-3-3S7 3.3 7 5v8c-2.2 1.7-2.7 4.8-1 7s4.8 2.7 7 1s2.7-4.8 1-7c-.3-.4-.6-.7-1-1m-3-9c.6 0 1 .4 1 1v3H9V5c0-.6.4-1 1-1Z"/>`),
		g.Group(children),
	)
}