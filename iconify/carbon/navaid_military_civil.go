package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidMilitaryCivil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 24a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8Zm0-14a6 6 0 1 0 6 6a6.007 6.007 0 0 0-6-6Z"/><path fill="currentColor" d="M30 14h-2.18A12.011 12.011 0 0 0 18 4.18V2h-4v2.18A12.011 12.011 0 0 0 4.18 14H2v4h2.18A12.011 12.011 0 0 0 14 27.82V30h4v-2.18A12.011 12.011 0 0 0 27.82 18H30ZM16 26a10 10 0 1 1 10-10a10.011 10.011 0 0 1-10 10Z"/>`),
		g.Group(children),
	)
}