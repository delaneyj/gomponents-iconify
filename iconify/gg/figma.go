package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 2a3 3 0 0 0 0 6h7a3 3 0 1 0 0-6h-7Zm7 7a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-10 3a3 3 0 0 1 3-3h3v6h-3a3 3 0 0 1-3-3Zm3 4a3 3 0 1 0 3 3v-3h-3Z"/>`),
		g.Group(children),
	)
}