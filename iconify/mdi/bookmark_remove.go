package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3a2 2 0 0 1 2 2v16l-7-3l-7 3V5a2 2 0 0 1 2-2h10M8.17 8.58L10.59 11l-2.42 2.41l1.42 1.42L12 12.41l2.41 2.42l1.42-1.42L13.41 11l2.42-2.42l-1.42-1.41L12 9.58L9.59 7.17L8.17 8.58Z"/>`),
		g.Group(children),
	)
}