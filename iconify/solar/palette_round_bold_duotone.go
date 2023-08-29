package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteRoundBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 18a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M10 6v12a4 4 0 0 1-8 0V6a4 4 0 1 1 8 0Z" opacity=".4"/><path d="m9.248 20.336l3.974-3.975l5.838-6.09a4.042 4.042 0 0 0-5.776-5.655L10 7.9V18c0 .872-.279 1.679-.752 2.336Z" opacity=".7"/><path d="m13.222 16.362l-3.974 3.974A4 4 0 0 1 6 22h11.9a4 4 0 1 0 0-8h-2.414l-2.264 2.362Z"/></g>`),
		g.Group(children),
	)
}