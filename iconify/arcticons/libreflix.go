package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libreflix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41 5H7a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2Zm-9.86 20.1l-12.88 8.68a1.26 1.26 0 0 1-1.76-.28a1.24 1.24 0 0 1-.24-.72V15.26A1.25 1.25 0 0 1 17.51 14a1.27 1.27 0 0 1 .75.25L31.15 23a1.25 1.25 0 0 1 .36 1.73a1.19 1.19 0 0 1-.37.37Z"/>`),
		g.Group(children),
	)
}