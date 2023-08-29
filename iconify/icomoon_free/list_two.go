package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 1h10v2H6V1zm0 6h10v2H6V7zm0 6h10v2H6v-2zM0 2a2 2 0 1 1 3.999-.001A2 2 0 0 1 0 2zm0 6a2 2 0 1 1 3.999-.001A2 2 0 0 1 0 8zm0 6a2 2 0 1 1 3.999-.001A2 2 0 0 1 0 14z"/>`),
		g.Group(children),
	)
}