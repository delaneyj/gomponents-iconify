package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M66.56 91.07L22.25 128l44.31 36.93A4 4 0 0 1 64 172a3.94 3.94 0 0 1-2.56-.93l-48-40a4 4 0 0 1 0-6.14l48-40a4 4 0 0 1 5.12 6.14Zm176 33.86l-48-40a4 4 0 1 0-5.12 6.14L233.75 128l-44.31 36.93a4 4 0 1 0 5.12 6.14l48-40a4 4 0 0 0 0-6.14Zm-81.19-88.69a4 4 0 0 0-5.13 2.39l-64 176a4 4 0 0 0 2.39 5.13A4.12 4.12 0 0 0 96 220a4 4 0 0 0 3.76-2.63l64-176a4 4 0 0 0-2.39-5.13Z"/>`),
		g.Group(children),
	)
}