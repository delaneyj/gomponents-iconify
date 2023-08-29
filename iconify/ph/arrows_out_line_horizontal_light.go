package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutLineHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M134 40v176a6 6 0 0 1-12 0V40a6 6 0 0 1 12 0Zm-38 82H30.49l21.75-21.76a6 6 0 0 0-8.48-8.48l-32 32a6 6 0 0 0 0 8.48l32 32a6 6 0 0 0 8.48-8.48L30.49 134H96a6 6 0 0 0 0-12Zm148.24 1.76l-32-32a6 6 0 0 0-8.48 8.48L225.51 122H160a6 6 0 0 0 0 12h65.51l-21.75 21.76a6 6 0 1 0 8.48 8.48l32-32a6 6 0 0 0 0-8.48Z"/>`),
		g.Group(children),
	)
}