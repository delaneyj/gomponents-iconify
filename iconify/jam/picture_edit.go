package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 13V9a1 1 0 1 1-2 0V7a1 1 0 1 1 2 0V3a1 1 0 1 1-2 0a3 3 0 0 1 3-3a1 1 0 1 1 0 2h4a1 1 0 1 1 0-2h3a1 1 0 0 1 0 2h4a1 1 0 0 1 0-2h3a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3a1 1 0 0 1 2 0Zm16-4.497V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h3.504l4.39-7.322a3 3 0 0 1 4.69-.582L18 8.503Zm0 2.823l-3.828-3.814a1 1 0 0 0-1.563.195L8.836 14H17a1 1 0 0 0 1-1v-1.674ZM6 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}