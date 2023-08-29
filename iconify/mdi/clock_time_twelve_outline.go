package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTimeTwelveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20c4.42 0 8-3.58 8-8s-3.58-8-8-8s-8 3.58-8 8s3.58 8 8 8m0-18c5.5 0 10 4.5 10 10s-4.5 10-10 10C6.47 22 2 17.5 2 12S6.5 2 12 2m.5 11.03H11V7h1.5v6.03Z"/>`),
		g.Group(children),
	)
}