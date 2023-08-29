package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliderTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 18h7M3 18h2m0 0a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Zm15-6h1M3 12h7m3-6h8m-8 0a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0ZM3 6h1m12.5 8.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}