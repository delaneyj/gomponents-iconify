package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Error(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="m29.582 8.683l-.129.12L8.3 29.954a3.308 3.308 0 0 0-.547.688c-2.04-2.639-3.233-6-3.233-9.701c0-8.797 6.626-15.482 15.421-15.482c3.691 0 7.014 1.185 9.641 3.224zM10.937 33.704c.189-.117.388-.287.606-.507l21.151-21.151l.041-.04c1.74 2.518 2.746 5.602 2.746 8.994c0 8.785-6.696 15.541-15.481 15.541c-3.432 0-6.546-1.035-9.063-2.837zM.5 21C.5 31.775 9.235 40.5 20 40.5c10.767 0 19.501-8.725 19.501-19.5s-8.734-19.5-19.5-19.5S.5 10.225.5 21z"/>`),
		g.Group(children),
	)
}