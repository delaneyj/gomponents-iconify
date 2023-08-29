package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CemeteryJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11.5 11h-3V2.5a.945.945 0 0 0-1-1a.945.945 0 0 0-1 1V11h-3a1 1 0 0 0 0 2h8a.945.945 0 0 0 1-1a1.002 1.002 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}