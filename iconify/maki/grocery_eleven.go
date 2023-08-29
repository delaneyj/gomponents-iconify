package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroceryEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.75 1.5c-.004 0-.318-.006-.63.188C8.803 1.881 8.5 2.3 8.5 3H1l.75 3.5c.107.5.75.5.75.5h6s-.003.247-.152.496S7.93 8 7.25 8H2c-.338-.005-.338.505 0 .5h5.25c.82 0 1.302-.37 1.527-.746C9.003 7.378 9 7 9 7V3c0-.567.196-.772.38-.887C9.567 2 9.747 2 9.747 2h.504c.338.005.338-.505 0-.5h-.5zm-2.5 7a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5zm-4 0a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}