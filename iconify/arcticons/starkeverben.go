package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Starkeverben(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.5 25.82H24a2.54 2.54 0 1 1 0-5.08h9.5c1.41 0-.39 1.14-.39 2.54s1.8 2.54.39 2.54Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 15.32a8 8 0 1 0 0 15.93h0a7.13 7.13 0 1 1 0 14.25h0a21.49 21.49 0 1 1 13.52-4.78"/>`),
		g.Group(children),
	)
}