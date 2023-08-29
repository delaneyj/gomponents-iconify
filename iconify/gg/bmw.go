package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bmw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0ZM5 12a7 7 0 0 0 7 7v-7h7a7 7 0 0 0-7-7v7H5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}