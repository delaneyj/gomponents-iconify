package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3A10 10 0 0 0 2 13v7a1 1 0 0 0 1 1h3a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3H4a8 8 0 0 1 16 0h-2a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h3a1 1 0 0 0 1-1v-7A10 10 0 0 0 12 3ZM6 15a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H4v-4Zm14 4h-2a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h2Z"/>`),
		g.Group(children),
	)
}