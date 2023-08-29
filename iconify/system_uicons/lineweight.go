package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lineweight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor"><rect width="14" height="2" x="3.5" y="6.5" fill="currentColor" rx="1"/><path fill="currentColor" d="M3.5 11.5h14v1h-14z"/><path stroke-linecap="round" stroke-linejoin="round" d="M3.5 15.5h13.981"/></g>`),
		g.Group(children),
	)
}