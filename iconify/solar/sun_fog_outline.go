package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunFogOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8 22h8M5 19h14M2 16h20"/><path d="M12 6a6 6 0 0 0-4.5 9.969h9A6 6 0 0 0 12 6Z"/><path stroke-linecap="round" d="M12 2v1m10 9h-1M3 12H2m17.07-7.07l-.392.393M5.322 5.322l-.393-.393"/></g>`),
		g.Group(children),
	)
}