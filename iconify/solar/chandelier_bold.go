package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChandelierBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.25 4A.75.75 0 0 1 9 3.25h6a.75.75 0 0 1 0 1.5h-2.25V16.5a2.75 2.75 0 1 0 5.5 0v-.595A3.001 3.001 0 0 1 16 13v-1.8a1.2 1.2 0 0 1 1.2-1.2h3.6a1.2 1.2 0 0 1 1.2 1.2V13a3.001 3.001 0 0 1-2.25 2.905v.595A4.25 4.25 0 0 1 12 18.912A4.25 4.25 0 0 1 4.25 16.5v-.595A3.001 3.001 0 0 1 2 13v-2.143c0-.473.384-.857.857-.857h4.286c.473 0 .857.384.857.857V13a3.001 3.001 0 0 1-2.25 2.905v.595a2.75 2.75 0 1 0 5.5 0V4.75H9A.75.75 0 0 1 8.25 4Z"/>`),
		g.Group(children),
	)
}