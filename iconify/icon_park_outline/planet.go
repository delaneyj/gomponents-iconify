package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Planet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M24 40c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Z"/><path stroke-linecap="round" d="M37.564 15.51c4.22.368 7.115 1.662 7.686 3.796c1.144 4.268-7.443 10.277-19.18 13.421c-11.736 3.145-22.177 2.235-23.32-2.033c-.6-2.236 1.472-4.95 5.269-7.48"/></g>`),
		g.Group(children),
	)
}