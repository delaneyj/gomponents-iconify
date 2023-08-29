package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230 128a6 6 0 0 1-6 6H86.49l61.75 61.76a6 6 0 1 1-8.48 8.48l-72-72a6 6 0 0 1 0-8.48l72-72a6 6 0 0 1 8.48 8.48L86.49 122H224a6 6 0 0 1 6 6ZM40 34a6 6 0 0 0-6 6v176a6 6 0 0 0 12 0V40a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}