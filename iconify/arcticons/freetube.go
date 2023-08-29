package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freetube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 42.5h3v-35a2 2 0 0 0-2-2h-7a2 2 0 0 0-2 2v27a8 8 0 0 0 8 8Zm3-35v7a2 2 0 0 0 2 2h16a8 8 0 0 0 8-8v-3h-24a2 2 0 0 0-2 2Zm13.907 15.491l-12.484-5.827a1 1 0 0 0-1.423.906v11.655a1 1 0 0 0 1.423.906l12.485-5.827a1 1 0 0 0 0-1.813Z"/>`),
		g.Group(children),
	)
}