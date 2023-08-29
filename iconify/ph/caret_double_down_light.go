package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleDownLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212.24 123.76a6 6 0 0 1 0 8.48l-80 80a6 6 0 0 1-8.48 0l-80-80a6 6 0 1 1 8.48-8.48L128 199.51l75.76-75.75a6 6 0 0 1 8.48 0Zm-88.48 8.48a6 6 0 0 0 8.48 0l80-80a6 6 0 0 0-8.48-8.48L128 119.51L52.24 43.76a6 6 0 0 0-8.48 8.48Z"/>`),
		g.Group(children),
	)
}