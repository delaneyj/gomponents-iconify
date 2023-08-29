package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowleftTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.87 9.814L15.685 16l6.187 6.188l-3.535 3.537L8.612 16l9.723-9.724z"/>`),
		g.Group(children),
	)
}