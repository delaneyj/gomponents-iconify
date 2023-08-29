package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChandelierBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11.25 4.75V16.5a2.75 2.75 0 1 1-5.5 0v-.595h-1.5v.595A4.25 4.25 0 0 0 12 18.912a4.25 4.25 0 0 0 7.75-2.412v-.595h-1.5v.595a2.75 2.75 0 1 1-5.5 0V4.75h-1.5Z" clip-rule="evenodd" opacity=".5"/><path d="M9 3.25a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5H9ZM16 13a3 3 0 0 0 2.25 2.905h1.5A3.001 3.001 0 0 0 22 13v-1.8a1.2 1.2 0 0 0-1.2-1.2h-3.6a1.2 1.2 0 0 0-1.2 1.2V13ZM2 13a3 3 0 0 0 2.25 2.905h1.5A3.001 3.001 0 0 0 8 13v-2.143A.857.857 0 0 0 7.143 10H2.857a.857.857 0 0 0-.857.857V13Z"/></g>`),
		g.Group(children),
	)
}