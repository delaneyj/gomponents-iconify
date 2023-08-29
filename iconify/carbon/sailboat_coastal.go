package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SailboatCoastal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24.78 26l1.19-4.758A1 1 0 0 0 25 20h-5v-3h5a1 1 0 0 0 .908-1.419l-6-13a1 1 0 0 0-1.702-.19l-9.998 13A1 1 0 0 0 9 17h9v3H7a1 1 0 0 0-.97 1.242L7.22 26H2v2h28v-2ZM20 7.553L23.437 15H20ZM11.03 15L18 5.94V15Zm11.69 11H9.28l-1-4h15.44Z"/>`),
		g.Group(children),
	)
}