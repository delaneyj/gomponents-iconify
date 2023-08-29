package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleZeroFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2zm0 5a3 3 0 0 0-2.995 2.824L9 10v4l.005.176a3 3 0 0 0 5.99 0L15 14v-4l-.005-.176A3 3 0 0 0 12 7zm0 2a1 1 0 0 1 .993.883L13 10v4l-.007.117a1 1 0 0 1-1.986 0L11 14v-4l.007-.117A1 1 0 0 1 12 9z"/></g>`),
		g.Group(children),
	)
}