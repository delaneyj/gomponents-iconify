package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OperaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 12.5a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3m0 10a3 3 0 0 0 3-3v-4a3 3 0 0 0-3-3m0 10h.585c1.907 0 3.78-.518 5.415-1.5m-6-8.5h.585c1.907 0 3.78.518 5.415 1.5m-6 10.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		g.Group(children),
	)
}