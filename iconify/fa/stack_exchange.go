package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackExchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1259 997v66q0 85-57.5 144.5T1063 1267h-57l-260 269v-269H217q-81 0-138.5-59.5T21 1063v-66h1238zm0-326v255H21V671h1238zm0-328v255H21V343h1238zm0-140v67H21v-67q0-84 57.5-143.5T217 0h846q81 0 138.5 59.5T1259 203z"/>`),
		g.Group(children),
	)
}