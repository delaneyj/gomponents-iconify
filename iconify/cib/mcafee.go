package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mcafee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 6.459L2.145.042v25.5L16 31.959l13.855-6.417V.042zm8.208 15.458l-8.172 3.78l-8.167-3.78V8.937l8.167 3.781l8.172-3.781z"/>`),
		g.Group(children),
	)
}