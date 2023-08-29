package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 18.17l-2.59-2.58L7 17l5 5l5-5l-1.41-1.41L13 18.17V2h-2v16.17Z"/>`),
		g.Group(children),
	)
}