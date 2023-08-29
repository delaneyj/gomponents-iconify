package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M222 1000H0V667h222v333zm371 0H370V333h223v667zm148 0V0h222v1000H741z"/>`),
		g.Group(children),
	)
}