package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zeronet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l-13.73 7.92v15.86l6.44-3.75v-8.39L24 11.93l7.19 4.15l-20.92 12.2v7.3L24 43.5l13.73-7.93V19.72l-6.44 3.75v8.39L24 36.07l-7.19-4.15l20.92-12.2v-7.29Z"/>`),
		g.Group(children),
	)
}