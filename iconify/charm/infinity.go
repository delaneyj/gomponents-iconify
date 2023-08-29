package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 5c2.5 1 3.5 5 6 6s3.25-1.25 3.25-3S13.5 4 11 5s-3.5 5-6 6s-3.25-1.25-3.25-3S2.5 4 5 5z"/>`),
		g.Group(children),
	)
}