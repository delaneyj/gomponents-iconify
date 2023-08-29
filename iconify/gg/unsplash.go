package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unsplash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 4.5H9v4h6v-4Zm-11 6h5v4h6v-4h5v9H4v-9Z"/>`),
		g.Group(children),
	)
}