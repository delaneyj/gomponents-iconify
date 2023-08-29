package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropInvert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.136L5.636 7.5a9 9 0 0 0 7.227 15.323A9 9 0 0 0 18.364 7.5L12 1.136ZM7.05 8.914L12 3.964v16.9a7 7 0 0 1-4.95-11.95Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}