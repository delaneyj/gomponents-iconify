package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basecamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M2 16.25s2-12.5 10-12.5s10 12.5 10 12.5s-2 4-10 4s-10-4-10-4m1.35-.6S4.3 19 12 19c5 0 8-1.2 8.65-3.15c.65-1.95-5-8.25-6-8.25S11.2 12 10.45 12c-2 0-1.55-2-3.3-2c-1.75 0-3.8 5.65-3.8 5.65z" fill="currentColor"/>`),
		g.Group(children),
	)
}