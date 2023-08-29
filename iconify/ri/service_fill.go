package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.121 10.48a1 1 0 0 0-1.414 0l-.707.707a2 2 0 0 1-2.828-2.829l5.63-5.632a6.5 6.5 0 0 1 6.377 10.568l-2.108 2.135l-4.95-4.95ZM3.161 4.47a6.503 6.503 0 0 1 8.009-.938L7.757 6.944a4 4 0 0 0 5.513 5.795l.144-.138l4.243 4.243l-4.243 4.242a2 2 0 0 1-2.828 0L3.16 13.662a6.5 6.5 0 0 1 0-9.193Z"/>`),
		g.Group(children),
	)
}