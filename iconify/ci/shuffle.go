package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 20l3-3m0 0l-3-3m3 3h-4a5 5 0 0 1-5-5a5 5 0 0 0-5-5H3m15-3l3 3m0 0l-3 3m3-3h-4a4.978 4.978 0 0 0-3 1M3 17h4a4.978 4.978 0 0 0 3-1"/>`),
		g.Group(children),
	)
}