package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18m-4-4v3l-5-3l-5 3V7m1.178-2.818c.252-.113.53-.176.822-.176h6a2 2 0 0 1 2 2v7"/>`),
		g.Group(children),
	)
}