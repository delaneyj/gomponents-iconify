package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 4h12"/><path d="M24 44a6 6 0 0 0 6-6V10H18v28a6 6 0 0 0 6 6Z" clip-rule="evenodd"/><path d="M24 27v1m0-10v3m-5 14h11"/></g>`),
		g.Group(children),
	)
}