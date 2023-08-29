package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendDownLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230 56a102.12 102.12 0 0 1-102 102H46.49l37.75 37.76a6 6 0 1 1-8.48 8.48l-48-48a6 6 0 0 1 0-8.48l48-48a6 6 0 0 1 8.48 8.48L46.49 146H128a90.1 90.1 0 0 0 90-90a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}