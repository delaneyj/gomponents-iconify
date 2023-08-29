package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretDownSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m5.5 4.793l2 2l2-2l.707.707L7.5 8.207L4.793 5.5l.707-.707Zm0 3l2 2l2-2l.707.707L7.5 11.207L4.793 8.5l.707-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}