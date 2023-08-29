package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JpgSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 8h.5a.5.5 0 0 0 0-1H7v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM4 7H2V6h3v5H2V9h1v1h1V7Zm2-1h1.5a1.5 1.5 0 1 1 0 3H7v2H6V6Zm4 0h3v1h-2v3h1V8.5h1V11h-3V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}