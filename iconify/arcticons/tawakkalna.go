package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tawakkalna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.942 39.992a7.402 7.402 0 0 1-5.251-2.176l-7.016-7.015a7.426 7.426 0 0 1 10.503-10.502l1.764 1.764l11.88-11.88a7.427 7.427 0 0 1 10.503 10.503L24.193 37.817a7.406 7.406 0 0 1-5.251 2.175Z"/>`),
		g.Group(children),
	)
}