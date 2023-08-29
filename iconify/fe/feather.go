package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFeather0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFeather1" fill="currentColor"><path id="feFeather2" d="M5.993 17.877c.004.04.007.081.007.123v3a1 1 0 0 1-2 0v-3C4 9.163 11.163 2 20 2c0 8.162-6.111 14.896-14.007 15.877Zm.174-2.044A14.01 14.01 0 0 0 17.833 4.167A14.01 14.01 0 0 0 6.167 15.833Z"/></g></g>`),
		g.Group(children),
	)
}