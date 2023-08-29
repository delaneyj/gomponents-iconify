package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M14 2a1 1 0 0 1 .993.883L15 3v2h3a1 1 0 0 1 .993.883L19 6v5a1 1 0 0 1-.883.993L18 12h-3v8h1a1 1 0 0 1 .117 1.993L16 22h-4a1 1 0 0 1-.117-1.993L12 20h1v-8H8a1 1 0 0 1-.694-.28l-.087-.095l-2-2.5a1 1 0 0 1-.072-1.147l.072-.103l2-2.5a1 1 0 0 1 .652-.367L8 5h5V3a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}