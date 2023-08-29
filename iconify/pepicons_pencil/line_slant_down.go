package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSlantDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.454.454a.5.5 0 0 1 .707 0l18.385 18.385a.5.5 0 0 1-.707.707L.454 1.16a.5.5 0 0 1 0-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}