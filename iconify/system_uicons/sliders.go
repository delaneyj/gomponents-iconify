package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sliders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 9V2.5m0 16V14"/><circle cx="14.5" cy="11.5" r="2.5"/><path d="M6.5 5V2.5m0 16V10"/><circle cx="6.5" cy="7.5" r="2.5"/></g>`),
		g.Group(children),
	)
}