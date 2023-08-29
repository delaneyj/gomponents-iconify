package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SailboatOffshore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.97 25h2.06l-.75-3h15.44l-.75 3h2.06l.94-3.758A1 1 0 0 0 25 20h-5v-3h5a1 1 0 0 0 .908-1.419l-6-13a1 1 0 0 0-1.702-.19l-9.998 13A1 1 0 0 0 9 17h9v3H7a1 1 0 0 0-.97 1.242ZM20 7.553L23.437 15H20ZM11.03 15L18 5.94V15Z"/><path fill="currentColor" d="M26.906 25.751A2.998 2.998 0 0 1 24 28a3.438 3.438 0 0 1-3.051-2.316A1.007 1.007 0 0 0 20 25a.99.99 0 0 0-.947.679A3.44 3.44 0 0 1 16 28a3.438 3.438 0 0 1-3.051-2.316A1.007 1.007 0 0 0 12 25a.971.971 0 0 0-.947.679A3.44 3.44 0 0 1 8 28a2.998 2.998 0 0 1-2.906-2.249l-1.936.498A4.997 4.997 0 0 0 8 30a4.932 4.932 0 0 0 4-1.987a5.02 5.02 0 0 0 8 0A4.932 4.932 0 0 0 24 30a4.997 4.997 0 0 0 4.842-3.751Z"/>`),
		g.Group(children),
	)
}