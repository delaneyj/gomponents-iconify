package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoolBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 64V40a20 20 0 0 0-20-20H72a20 20 0 0 0-20 20v24a20 20 0 0 0 20 20h2L52.15 222.13a12 12 0 1 0 23.7 3.74L83.1 180h89.8l7.25 45.87a12 12 0 1 0 23.7-3.74L182 84h2a20 20 0 0 0 20-20ZM76 44h104v16H76Zm93.11 112H86.89l11.36-72h59.5Z"/>`),
		g.Group(children),
	)
}