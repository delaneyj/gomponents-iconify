package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackInTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 1.799c-4.445 0-8.061 3.562-8.169 7.996V10H.459l3.594 3.894L7.547 10H4.875v-.205C4.982 6.492 7.683 3.85 11 3.85c3.386 0 6.131 2.754 6.131 6.15c0 3.396-2.745 6.15-6.131 6.15a6.099 6.099 0 0 1-3.627-1.193l-1.406 1.504A8.13 8.13 0 0 0 11 18.199c4.515 0 8.174-3.67 8.174-8.199S15.515 1.799 11 1.799zM10 5v5a1.01 1.01 0 0 0 .293.707l3.2 3.2c.283-.183.55-.389.787-.628L12 11V5h-2z"/>`),
		g.Group(children),
	)
}