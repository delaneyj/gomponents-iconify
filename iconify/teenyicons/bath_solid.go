package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BathSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2 3.5A2.5 2.5 0 0 1 4.5 1H6v1h1V0H4.5A3.5 3.5 0 0 0 1 3.5V7H0v1h1v1.5a3.5 3.5 0 0 0 3 3.465V15h1v-2h5v2h1v-2.035A3.501 3.501 0 0 0 14 9.5V8h1V7H2V3.5Z"/><path fill="currentColor" d="M8 4H5V3h3v1Z"/>`),
		g.Group(children),
	)
}