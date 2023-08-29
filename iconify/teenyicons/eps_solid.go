package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EpsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 8h1V7H7v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM5 6H2v5h3v-1H3V9h2V8H3V7h2V6Zm1 0h3v3H7v2H6V6Zm4 0h3v1h-2v1h2v3h-3v-1h2V9h-2V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}