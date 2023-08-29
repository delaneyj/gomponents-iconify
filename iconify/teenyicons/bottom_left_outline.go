package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 13.5H1a.5.5 0 0 0 .5.5v-.5Zm0 .5H7v-1H1.5v1Zm.5-.5V8H1v5.5h1Zm-.146.354l12-12l-.708-.708l-12 12l.708.708Z"/>`),
		g.Group(children),
	)
}