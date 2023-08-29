package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7.5 0L1 3.445V4h13v-.555L7.5 0zM2 5v5l-1 1.555V13h13v-1.445L13 10V5H2zm2.615 1a.625.625 0 0 1 .451.184L7.5 8.617l2.434-2.433A.625.625 0 0 1 11 6.625v4.242a.625.625 0 1 1-1.25 0V8.133L7.941 9.94a.625.625 0 0 1-.882 0L5.25 8.133v2.734a.625.625 0 1 1-1.25 0V6.625c0-.341.274-.62.615-.625z" fill="currentColor"/>`),
		g.Group(children),
	)
}