package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Random(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M6 0v1h-.5c-.35 0-.56.1-.78.38L3.31 3.16L1.78 1.38C1.56 1.12 1.34 1 1 1H0v1h1c-.05 0 .01.04.03.03l1.63 1.91L1 6H0v1h1c.35 0 .56-.1.78-.38l1.53-1.91l1.66 1.91c.22.26.44.38.78.38H6v1l2-1.5L6 5v1h-.22c-.01-.01-.05-.04-.06-.03L3.97 3.91L5.5 2H6v1l2-1.5L6 0z"/>`),
		g.Group(children),
	)
}