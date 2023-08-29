package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignRightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M10 2a1 1 0 0 1 .993.883L11 3v2h5a1 1 0 0 1 .694.28l.087.095l2 2.5a1 1 0 0 1 .072 1.147l-.072.103l-2 2.5a1 1 0 0 1-.652.367L16 12h-5v8h1a1 1 0 0 1 .117 1.993L12 22H8a1 1 0 0 1-.117-1.993L8 20h1v-8H6a1 1 0 0 1-.993-.883L5 11V6a1 1 0 0 1 .883-.993L6 5h3V3a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}