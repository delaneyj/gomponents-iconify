package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vidmate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4.5 13.247l3.207-1.862h8.89l13.51 25.265L43.5 11.35h-8.89l-8.938 16.885"/><path d="M8.711 13.263L21.218 36.65h8.89M16.597 11.385h0l-3.207 1.862H4.5"/></g>`),
		g.Group(children),
	)
}