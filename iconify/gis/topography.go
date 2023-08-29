package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Topography(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M36.771 7.082a3.274 3.274 0 0 0-3.273 3.273v4.582H15.063v45.38h18.435v4.581a3.274 3.274 0 0 0 3.273 3.274h3.575v4.746h3.588L39.143 100h5.67l4.79-27.082h.793L55.19 100h5.67l-4.79-27.082h3.585v-4.746h3.575a3.274 3.274 0 0 0 3.273-3.274v-4.582h18.436V14.937H66.501v-4.582a3.274 3.274 0 0 0-3.273-3.273H36.77zm3.272 6.547h19.914v37.523H40.043V13.63zm9.885 6.816a6.708 6.708 0 0 0-6.637 6.71a6.709 6.709 0 1 0 6.637-6.71z" color="currentColor"/>`),
		g.Group(children),
	)
}