package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func U(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 10a2 2 0 0 1 2 2v16a6 6 0 0 0 12 0V12a2 2 0 1 1 4 0v16c0 5.523-4.477 10-10 10s-10-4.477-10-10V12a2 2 0 0 1 2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}