package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 9h2.31l.32-3h2l-.32 3h2l.32-3h2l-.32 3H15v2h-1.9l-.2 2H15v2h-2.31l-.32 3h-2l.32-3h-2l-.32 3h-2l.32-3H5v-2h1.9l.2-2H5V9m4.1 2l-.2 2h2l.2-2M19 6h-2v8h2m0 2h-2v2h2Z"/>`),
		g.Group(children),
	)
}