package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pictogrammers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 0h14v1h1v1h1v18h-1v1h-1v1H4v-1H3v-1H2V2h1V1h1V0m0 18v1h1v1h12v-1h1v-1H4M17 2H5v1H4v12h1v1h12v-1h1V3h-1V2m-4 2v1h1v1h1v2h-1v1h-1v1h-3v4H8V4h5m-1 2h-2v2h2V6Z"/>`),
		g.Group(children),
	)
}