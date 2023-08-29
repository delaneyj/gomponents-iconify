package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func One(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M27.765 10.152A2 2 0 0 1 29 12v24a2 2 0 0 1-4 0V16.828l-1.586 1.586a2 2 0 0 1-2.828-2.828l5-5a2 2 0 0 1 2.18-.434Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}