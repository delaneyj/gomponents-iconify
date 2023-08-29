package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 9h5.485a1 1 0 0 1 .814.419l9.402 13.162a1 1 0 0 0 .814.419H27M21 9h6"/>`),
		g.Group(children),
	)
}