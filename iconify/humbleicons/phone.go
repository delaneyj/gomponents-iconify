package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.515 4.621L9 4l2 3.5L9.5 9c1 2 3.5 4.5 5.5 5.5l1.5-1.5l3.5 2l-.621 2.485c-.223.89-1.029 1.534-1.928 1.351c-5.213-1.06-11.228-7.074-12.287-12.287c-.183-.9.46-1.705 1.35-1.928Z"/>`),
		g.Group(children),
	)
}