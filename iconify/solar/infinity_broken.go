package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfinityBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M10 8a5 5 0 0 0-6 8m10 0a5 5 0 0 0 6-8M7 17c2.761 0 3.5-2 5-5s2.239-5 5-5"/>`),
		g.Group(children),
	)
}