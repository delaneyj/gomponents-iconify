package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageBreakSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V8H1V1.5ZM1 11h13v2.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V11ZM0 8.993h3v1H0v-1Zm4 0h3v1H4v-1Zm7 0H8v1h3v-1Zm1 0h3v1h-3v-1Z"/>`),
		g.Group(children),
	)
}