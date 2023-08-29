package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRadiusSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 2A3.5 3.5 0 0 0 2 5.5V8H1V5.5A4.5 4.5 0 0 1 5.5 1H8v1H5.5ZM11 2h-1V1h1v1Zm3 0h-1V1h1v1Zm0 3h-1V4h1v1Zm0 3h-1V7h1v1ZM2 11H1v-1h1v1Zm12 0h-1v-1h1v1ZM2 14H1v-1h1v1Zm3 0H4v-1h1v1Zm3 0H7v-1h1v1Zm3 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}