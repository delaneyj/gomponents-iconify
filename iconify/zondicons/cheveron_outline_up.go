package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheveronOutlineUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 10a10 10 0 1 1 20 0a10 10 0 0 1-20 0zm10 8a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm.7-10.54L14.25 11l-1.41 1.41L10 9.6l-2.83 2.8L5.76 11L10 6.76l.7.7z"/>`),
		g.Group(children),
	)
}