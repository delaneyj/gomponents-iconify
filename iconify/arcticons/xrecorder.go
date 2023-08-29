package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xrecorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.203 39.203A21.433 21.433 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5a21.433 21.433 0 0 1 15.203 6.297"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m44 15.588l-7.678 4.433V16.65a2 2 0 0 0-2-2H13.678a2 2 0 0 0-2 2v14.7a2 2 0 0 0 2 2h20.644a2 2 0 0 0 2-2v-3.37L44 32.412a1 1 0 0 0 1.5-.867v-15.09a1 1 0 0 0-1.5-.866Z"/><circle cx="24" cy="24" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}