package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartIncreasing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20c1.427 0 2.62.996 2.925 2.33L14.97 19.327a.5.5 0 0 1-.699.01l-2.148-2.04a2.2 2.2 0 0 0-3.06.03L3 23.31V6Zm.002 20.102l7.234-7.168a.5.5 0 0 1 .694-.01l2.15 2.018a2.2 2.2 0 0 0 3.052-.04L29 8.176V26a3 3 0 0 1-3 3H6a3 3 0 0 1-2.998-2.898Z"/>`),
		g.Group(children),
	)
}