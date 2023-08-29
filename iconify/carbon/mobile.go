package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 4H10a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h12a2.003 2.003 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2Zm0 2v2H10V6ZM10 28V10h12v18Z"/>`),
		g.Group(children),
	)
}