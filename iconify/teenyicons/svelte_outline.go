package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SvelteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m9.625 8.357l-5.088 3.18m2.968-1.855a3.5 3.5 0 0 1-3.71-5.937l4.241-2.65A3.5 3.5 0 0 1 12.405 6.5M7.536 5.296a3.5 3.5 0 0 1 3.71 5.936l-4.24 2.65A3.5 3.5 0 0 1 2.614 8.5m2.8-1.88l5.09-3.179"/>`),
		g.Group(children),
	)
}