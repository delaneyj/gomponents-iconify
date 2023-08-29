package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 11v10h10V11Zm8 8h-6v-6h6Z"/><path fill="currentColor" d="M30 13v-2h-4V8a2 2 0 0 0-2-2h-3V2h-2v4h-6V2h-2v4H8a2 2 0 0 0-2 2v3H2v2h4v6H2v2h4v3a2 2 0 0 0 2 2h3v4h2v-4h6v4h2v-4h3a2 2 0 0 0 2-2v-3h4v-2h-4v-6Zm-6 11H8V8h16Z"/>`),
		g.Group(children),
	)
}