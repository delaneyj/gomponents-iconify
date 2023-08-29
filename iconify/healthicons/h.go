package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func H(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 10a2 2 0 0 1 2 2v10h12V12a2 2 0 1 1 4 0v24a2 2 0 1 1-4 0V26H18v10a2 2 0 1 1-4 0V12a2 2 0 0 1 2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}