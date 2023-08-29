package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndustryEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10 1v8H1V6l2.11-1.78c.32-.32.89-.31.89.14V6l2.13-1.86A.5.5 0 0 1 7 4.5V8h2V1h1z" fill="currentColor"/>`),
		g.Group(children),
	)
}