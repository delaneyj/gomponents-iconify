package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaPlayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.998 7.002l.085.078L14.134 12l-5.096 4.964L9 17l-.002-9.998M9 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2c.543 0 1.033-.218 1.393-.568L17 12l-6.604-6.433A2.008 2.008 0 0 0 9 5z"/>`),
		g.Group(children),
	)
}