package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216.49 119.51a12 12 0 0 1 0 17l-80 80a12 12 0 0 1-17 0l-80-80a12 12 0 1 1 17-17L128 191l71.51-71.52a12 12 0 0 1 16.98.03Zm-97 17a12 12 0 0 0 17 0l80-80a12 12 0 0 0-17-17L128 111L56.49 39.51a12 12 0 0 0-17 17Z"/>`),
		g.Group(children),
	)
}