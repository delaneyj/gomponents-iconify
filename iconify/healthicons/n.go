package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func N(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 17.524V36a2 2 0 1 1-4 0V12a2 2 0 0 1 3.536-1.28L32 30.476V12a2 2 0 1 1 4 0v24a2 2 0 0 1-3.536 1.28L16 17.524Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}