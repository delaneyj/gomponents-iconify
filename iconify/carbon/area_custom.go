package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaCustom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 6a3.992 3.992 0 0 0-7.977-.224L9.586 8.263A3.99 3.99 0 1 0 5 13.858v8.284A3.991 3.991 0 1 0 9.858 27h8.284a3.991 3.991 0 1 0 5.595-4.586l2.487-12.437A3.994 3.994 0 0 0 30 6Zm-4-2a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2ZM4 10a2 2 0 1 1 2 2a2.002 2.002 0 0 1-2-2Zm2 18a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm12.142-3H9.858A3.994 3.994 0 0 0 7 22.142v-8.284a3.987 3.987 0 0 0 2.977-3.634l12.437-2.487a4.005 4.005 0 0 0 1.849 1.85l-2.487 12.436A3.987 3.987 0 0 0 18.142 25ZM22 28a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}