package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSlantUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.454 19.546a.5.5 0 0 1 0-.707L18.839.454a.5.5 0 1 1 .707.707L1.16 19.546a.5.5 0 0 1-.707 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}