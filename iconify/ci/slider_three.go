package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliderThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 18h11M3 18h3m0 0v2m0-2v-2m14-4h1M3 12h13m0 0v2m0-2v-2m-2-4h7M3 6h7m0 0v2m0-2V4"/>`),
		g.Group(children),
	)
}