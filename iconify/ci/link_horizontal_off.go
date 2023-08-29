package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkHorizontalOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h4m0 0L8 8m4 4l8 8M15 8h2a4 4 0 0 1 2.645 7M9 16H7a4 4 0 0 1 0-8h1m0 0L4 4"/>`),
		g.Group(children),
	)
}