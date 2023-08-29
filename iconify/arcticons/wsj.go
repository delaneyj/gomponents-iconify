package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wsj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.5 18.5l2.75 11l2.75-11l2.75 11l2.75-11m3.744 9.762a3.278 3.278 0 0 0 2.75 1.238h1.65a2.75 2.75 0 1 0 0-5.5h-1.788a2.75 2.75 0 0 1 0-5.5h1.65a2.952 2.952 0 0 1 2.75 1.237M39.5 18.5v8.25a2.75 2.75 0 0 1-5.5 0v-.962"/>`),
		g.Group(children),
	)
}