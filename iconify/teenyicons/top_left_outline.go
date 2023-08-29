package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 1.5V1a.5.5 0 0 0-.5.5h.5Zm0 .5H7V1H1.5v1ZM1 1.5V7h1V1.5H1Zm.146.354l12 12l.707-.708l-12-12l-.707.708Z"/>`),
		g.Group(children),
	)
}