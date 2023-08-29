package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUDownLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230 112a62.07 62.07 0 0 1-62 62H46.49l37.75 37.76a6 6 0 1 1-8.48 8.48l-48-48a6 6 0 0 1 0-8.48l48-48a6 6 0 0 1 8.48 8.48L46.49 162H168a50 50 0 0 0 0-100H80a6 6 0 0 1 0-12h88a62.07 62.07 0 0 1 62 62Z"/>`),
		g.Group(children),
	)
}