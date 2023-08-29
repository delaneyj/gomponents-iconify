package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eufyhome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.192 14.546L25.171 5.96a2.478 2.478 0 0 0-2.341 0L6.808 14.546A2.478 2.478 0 0 0 5.5 16.73v23.125a2.478 2.478 0 0 0 2.478 2.478h32.044a2.478 2.478 0 0 0 2.478-2.478V16.73a2.478 2.478 0 0 0-1.308-2.184Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 26.11l6.436-6.436a9.103 9.103 0 1 0 2.667 6.436"/>`),
		g.Group(children),
	)
}