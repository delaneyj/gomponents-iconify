package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DangerousOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.25 21L3 15.75v-7.5L8.25 3h7.5L21 8.25v7.5L15.75 21h-7.5Zm.9-4.75L12 13.4l2.85 2.85l1.4-1.4L13.4 12l2.85-2.85l-1.4-1.4L12 10.6L9.15 7.75l-1.4 1.4L10.6 12l-2.85 2.85l1.4 1.4ZM9.1 19h5.8l4.1-4.1V9.1L14.9 5H9.1L5 9.1v5.8L9.1 19Zm2.9-7Z"/>`),
		g.Group(children),
	)
}