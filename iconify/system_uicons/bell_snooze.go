package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellSnooze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.5 7.585c0 1.907.518 3.78 1.5 5.415a1.65 1.65 0 0 1-1.416 2.5H5.415A1.65 1.65 0 0 1 4 13a10.526 10.526 0 0 0 1.5-5.415V6.5a4 4 0 0 1 4-4h2a3.98 3.98 0 0 1 2.178.645"/><path d="M10.5 5.5h2l-2 3h2m2-7h3l-3 4h3M13 17c-.667 1-1.5 1.5-2.5 1.5S8.667 18 8 17"/></g>`),
		g.Group(children),
	)
}