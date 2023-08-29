package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m249 394l-1 1v-1L0 85Q113 0 248.5 0T497 85zM68 170q82-63 180.5-63T429 170L249 394l-1 1v-1z"/>`),
		g.Group(children),
	)
}