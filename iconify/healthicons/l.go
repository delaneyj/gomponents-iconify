package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func L(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a2 2 0 0 1 2 2v22h10a2 2 0 1 1 0 4H18a2 2 0 0 1-2-2V12a2 2 0 0 1 2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}