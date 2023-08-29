package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smalltalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="37" height="9.052" x="5.5" y="6.232" rx="2" ry="2"/><path d="M42.338 27.801V26a2 2 0 0 0-2-2H7.5a2 2 0 0 1-2-2v-1.801"/><rect width="37" height="9.052" x="5.5" y="32.716" rx="2" ry="2"/></g>`),
		g.Group(children),
	)
}