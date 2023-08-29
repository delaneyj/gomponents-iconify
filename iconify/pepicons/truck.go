package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M7.5 18.75a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5Zm9 0a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5Z"/><path fill-rule="evenodd" d="M18 4.5H8a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6.5a2 2 0 0 0-2-2ZM8 14V6.5h10V14H8Z" clip-rule="evenodd"/><path d="M10.25 8.75a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Zm0 2a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Zm0 2a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Z"/><path fill-rule="evenodd" d="M6 7.5H2.736a2 2 0 0 0-1.92 1.442L0 11.75V14a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V9.5a2 2 0 0 0-2-2ZM2 14v-1.966L2.736 9.5H6V14H2Z" clip-rule="evenodd"/><path d="M3.29 10.1H5a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5v-.85l.015-.122l.29-1.15a.5.5 0 0 1 .485-.378Z"/></g>`),
		g.Group(children),
	)
}