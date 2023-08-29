package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204.24 203.76a6 6 0 1 1-8.48 8.48l-80-80a6 6 0 0 1 0-8.48l80-80a6 6 0 0 1 8.48 8.48L128.49 128ZM48.49 128l75.75-75.76a6 6 0 0 0-8.48-8.48l-80 80a6 6 0 0 0 0 8.48l80 80a6 6 0 1 0 8.48-8.48Z"/>`),
		g.Group(children),
	)
}