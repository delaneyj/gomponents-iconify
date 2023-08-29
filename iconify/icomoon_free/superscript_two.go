package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperscriptTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m3.032 13l.9-3h4.137l.9 3h1.775l-3-10H4.256l-3 10h1.776zm2.4-8h1.137l.9 3H4.532l.9-3zM11 13l2.5-4l2.5 4h-5zm2.5-11h-1a.5.5 0 0 1 0-1h2a.5.5 0 0 0 0-1h-2a1.502 1.502 0 0 0-1.117 2.5c.275.307.674.5 1.117.5h1a.5.5 0 0 1 0 1h-2a.5.5 0 0 0 0 1h2a1.502 1.502 0 0 0 1.117-2.5A1.496 1.496 0 0 0 13.5 2z"/>`),
		g.Group(children),
	)
}