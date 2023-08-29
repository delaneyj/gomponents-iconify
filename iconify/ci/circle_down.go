package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm0 13l-5-5l1.41-1.41L11 13.17V7h2v6.17l2.59-2.58L17 12l-5 5Z"/>`),
		g.Group(children),
	)
}