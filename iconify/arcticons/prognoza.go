package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prognoza(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.894 24.166c1.388-9.906-3.846-19.15-12.686-20.706c-8.84-1.556-19.412 2.83-19.412 17.948"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.848 18.76c12.419 0 21.85 5.273 27.244 9.43c-1.928 10.104-11.164 18.162-23.014 17.233c-11.85-.93-19.517-11.502-19.517-20.31c0 0 4.235-6.352 15.287-6.352Z"/>`),
		g.Group(children),
	)
}