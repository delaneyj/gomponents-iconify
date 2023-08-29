package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleSevenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2zm2 5h-4l-.117.007a1 1 0 0 0-.876.876L9 8l.007.117a1 1 0 0 0 .876.876L10 9h2.718l-1.688 6.757l-.022.115a1 1 0 0 0 1.927.482l.035-.111l2-8l.021-.112a1 1 0 0 0-.878-1.125L14 7z"/></g>`),
		g.Group(children),
	)
}