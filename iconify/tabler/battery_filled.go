package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 6a3 3 0 0 1 2.995 2.824L20 9v.086l.052.019a1.5 1.5 0 0 1 .941 1.25L21 10.5v3a1.5 1.5 0 0 1-.948 1.395l-.052.018V15a3 3 0 0 1-2.824 2.995L17 18H6a3 3 0 0 1-2.995-2.824L3 15V9a3 3 0 0 1 2.824-2.995L6 6h11z"/></g>`),
		g.Group(children),
	)
}