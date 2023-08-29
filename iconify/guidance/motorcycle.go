package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Motorcycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 .5h3.5V1A2.5 2.5 0 0 0 10 3.5m10-3h-3.5V1A2.5 2.5 0 0 1 14 3.5m-4 3H4.5v.25l.365.638A20 20 0 0 1 7.5 17.311V19.5h3m3.5-13h5.5v.25l-.365.639A20 20 0 0 0 16.5 17.31v2.19h-3m-3-5V22a1.5 1.5 0 0 0 3 0v-7.5h-3Zm1.5-7a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}