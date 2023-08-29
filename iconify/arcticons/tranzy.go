package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tranzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.83 5.5v12.33H7.5a2 2 0 0 1-2-2V7.5a2 2 0 0 1 2-2h10.33Z"/><circle cx="24" cy="36.33" r="6.17"/><path d="M42.5 7.5v8.33a2 2 0 0 1-2 2h-8.33a2 2 0 0 0-2 2v8.34a2 2 0 0 1-2 2h-8.34a2 2 0 0 1-2-2V5.5H40.5a2 2 0 0 1 2 2Z"/></g>`),
		g.Group(children),
	)
}