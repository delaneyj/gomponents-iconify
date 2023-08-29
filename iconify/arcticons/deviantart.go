package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.363 43.5l4.071-7.768l12.64-.021v-9.754l-7.524-.009l7.524-14.347V4.5h-7.439l-4.071 7.793l-12.638.015v9.735l7.53-.003l-7.53 14.378V43.5h7.437z"/>`),
		g.Group(children),
	)
}