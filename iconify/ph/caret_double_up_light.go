package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleUpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212.24 203.76a6 6 0 1 1-8.48 8.48L128 136.49l-75.76 75.75a6 6 0 0 1-8.48-8.48l80-80a6 6 0 0 1 8.48 0Zm-160-71.52L128 56.49l75.76 75.75a6 6 0 0 0 8.48-8.48l-80-80a6 6 0 0 0-8.48 0l-80 80a6 6 0 0 0 8.48 8.48Z"/>`),
		g.Group(children),
	)
}