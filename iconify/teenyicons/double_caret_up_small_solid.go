package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretUpSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.793L10.207 6.5l-.707.707l-2-2l-2 2l-.707-.707L7.5 3.793Zm0 3L10.207 9.5l-.707.707l-2-2l-2 2l-.707-.707L7.5 6.793Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}