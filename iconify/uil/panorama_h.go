package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.54 5.16a1 1 0 0 0-1-.07A21.27 21.27 0 0 1 12 6.73a21.27 21.27 0 0 1-8.58-1.64a1 1 0 0 0-1 .07A1 1 0 0 0 2 6v12a1 1 0 0 0 .46.84a1 1 0 0 0 1 .07A21.27 21.27 0 0 1 12 17.27a21.27 21.27 0 0 1 8.58 1.64A1.06 1.06 0 0 0 21 19a1 1 0 0 0 .54-.16A1 1 0 0 0 22 18V6a1 1 0 0 0-.46-.84ZM20 16.52a24.77 24.77 0 0 0-8-1.25a24.77 24.77 0 0 0-8 1.25v-9a24.77 24.77 0 0 0 8 1.25a24.77 24.77 0 0 0 8-1.25Z"/>`),
		g.Group(children),
	)
}