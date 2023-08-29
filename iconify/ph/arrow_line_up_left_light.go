package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineUpLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 216a6 6 0 0 1-6 6H40a6 6 0 0 1 0-12h176a6 6 0 0 1 6 6ZM64 158a6 6 0 0 0 6-6V70.49l101.76 101.75a6 6 0 0 0 8.48-8.48L78.49 62H160a6 6 0 0 0 0-12H64a6 6 0 0 0-6 6v96a6 6 0 0 0 6 6Z"/>`),
		g.Group(children),
	)
}