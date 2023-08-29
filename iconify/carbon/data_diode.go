package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataDiode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15h-3V7h-2v8h-3V7h-2v7.826L7.524 7.148A1 1 0 0 0 6 8v7H2v2h4v7a1 1 0 0 0 1.524.852L20 17.174V25h2v-8h3v8h2v-8h3ZM8 22.21V9.79L18.092 16Z"/>`),
		g.Group(children),
	)
}