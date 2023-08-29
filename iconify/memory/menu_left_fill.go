package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 17h1V5h-1v1h-1v1h-1v1h-1v1H9v1H8v2h1v1h1v1h1v1h1v1h1"/>`),
		g.Group(children),
	)
}