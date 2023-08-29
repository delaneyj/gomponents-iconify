package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Implant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M21 9a3 3 0 1 1 6 0v30a3 3 0 1 1-6 0V9Z"/>`),
		g.Group(children),
	)
}