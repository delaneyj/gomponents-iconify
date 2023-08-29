package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckSquareOffsetLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 48v160a14 14 0 0 1-14 14h-72a6 6 0 0 1 0-12h72a2 2 0 0 0 2-2V48a2 2 0 0 0-2-2H48a2 2 0 0 0-2 2v96a6 6 0 0 1-12 0V48a14 14 0 0 1 14-14h160a14 14 0 0 1 14 14ZM115.76 155.76L64 207.51l-19.76-19.75a6 6 0 0 0-8.48 8.48l24 24a6 6 0 0 0 8.48 0l56-56a6 6 0 0 0-8.48-8.48Z"/>`),
		g.Group(children),
	)
}