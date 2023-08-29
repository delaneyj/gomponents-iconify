package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reuse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 2.5h5v5"/><path d="M8.5 2.5c-3.333 2.837-5 5.67-5 8.5s1 5.33 3 7.5m11 0h-5v-5"/><path d="M12.5 18.5c3.333-2.837 5-5.67 5-8.5s-1-5.33-3-7.5"/></g>`),
		g.Group(children),
	)
}