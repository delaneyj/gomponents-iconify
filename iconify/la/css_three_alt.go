package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThreeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6 4l2 21l8 3l8-3l2-21H6zm3.332 3h13.32l-.261 3l-5.696 3h5.428l-.512 6.008l.02-.008l-.276 3L16 24l-5.365-2l-.33-4h3.021l.156 2.033l2.518.871l2.521-.853l.346-4.051h-8.736l-.258-3l5.91-3H9.61l-.277-3z"/>`),
		g.Group(children),
	)
}