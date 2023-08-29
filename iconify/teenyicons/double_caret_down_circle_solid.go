package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretDownCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM4.793 5.5L7.5 8.207L10.207 5.5L9.5 4.793l-2 2l-2-2l-.707.707Zm0 3L7.5 11.207L10.207 8.5L9.5 7.793l-2 2l-2-2l-.707.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}