package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func M(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 17.524V36a2 2 0 1 1-4 0V12a2 2 0 0 1 3.536-1.28L24 20.876l8.464-10.156A2 2 0 0 1 36 12v24a2 2 0 1 1-4 0V17.524l-6.464 7.756a2 2 0 0 1-3.072 0L16 17.524Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}