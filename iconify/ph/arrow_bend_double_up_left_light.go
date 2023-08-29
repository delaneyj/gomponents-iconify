package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendDoubleUpLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M84.24 147.76a6 6 0 1 1-8.48 8.48l-48-48a6 6 0 0 1 0-8.48l48-48a6 6 0 0 1 8.48 8.48L40.49 104ZM128 98H94.49l37.75-37.76a6 6 0 0 0-8.48-8.48l-48 48a6 6 0 0 0 0 8.48l48 48a6 6 0 0 0 8.48-8.48L94.49 110H128a90.1 90.1 0 0 1 90 90a6 6 0 0 0 12 0A102.12 102.12 0 0 0 128 98Z"/>`),
		g.Group(children),
	)
}