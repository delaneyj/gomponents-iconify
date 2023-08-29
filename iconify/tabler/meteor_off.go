package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeteorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.75 5.761L13 3l-1 5l9-5l-5 9h5l-2.467 2.536m-1.983 2.04l-2.441 2.51A6.5 6.5 0 1 1 5.254 9.58l2.322-1.972"/><path d="M7 14.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0M3 3l18 18"/></g>`),
		g.Group(children),
	)
}