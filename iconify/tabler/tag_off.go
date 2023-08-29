package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.792 7.793a1 1 0 0 0 1.414 1.414"/><path d="M4.88 4.877A2.99 2.99 0 0 0 4 7v3.859c0 .537.213 1.052.593 1.432l8.116 8.116a2.025 2.025 0 0 0 2.864 0l2.416-2.416m2-2l.416-.416a2.025 2.025 0 0 0 0-2.864l-8.117-8.116a2.025 2.025 0 0 0-1.431-.593H7.998M3 3l18 18"/></g>`),
		g.Group(children),
	)
}