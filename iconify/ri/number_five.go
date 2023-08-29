package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2v2H9.3l-.676 6.445a6.5 6.5 0 1 1-2.93 7.133l1.94-.486A4.502 4.502 0 0 0 16.5 16a4.5 4.5 0 0 0-4.5-4.5c-2.022 0-3.277.639-3.96 1.53l-1.575-1.182L7.501 2H18Z"/>`),
		g.Group(children),
	)
}