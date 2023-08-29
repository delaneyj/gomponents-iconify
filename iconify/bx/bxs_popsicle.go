package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsPopsicle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 4a7 7 0 0 0-9.9 0l-7.82 7.82a1 1 0 0 0 0 1.41l3.54 3.54l-3.54 3.53l1.42 1.42l3.53-3.54l3.54 3.54a1 1 0 0 0 1.41 0L20 13.94A7 7 0 0 0 20 4zm-2.7 2.7a3.33 3.33 0 0 0-4.6 0l-1.06-1.06a4.76 4.76 0 0 1 6.72 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}