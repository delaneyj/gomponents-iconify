package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Restore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.22 0L0 1.22l4 4V17a2 2 0 0 0 2 2h8a2 2 0 0 0 2-1.8l2.8 2.8l1.2-1.22zM17 4V2h-3.5l-1-1h-5l-1 1h-.84l2 2zM8.66 5H16v7.34z"/>`),
		g.Group(children),
	)
}