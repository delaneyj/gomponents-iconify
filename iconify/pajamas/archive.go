package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4.5v-2h13v2h-13ZM1 6a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1v8a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V6Zm1.5 0v7.5h11V6h-11ZM4 8.75A.75.75 0 0 1 4.75 8h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 4 8.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}