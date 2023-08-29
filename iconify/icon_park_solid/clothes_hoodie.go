package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesHoodie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClothesHoodie0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37 22v15m-26 0v7h26v-7m-26 0H4V22c0-3 2-6.5 5-9s7-3 7-3l8 8M11 37V22m26 15h7V22c0-3-2-6.5-5-9s-7-3-7-3l-8 8m0 0v9"/><path fill="#fff" d="M9 13c3-2.5 7-3 7-3l8 8l8-8s4 .5 7 3l2-4.5L39 4H9L7 8.5L9 13Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClothesHoodie0)"/>`),
		g.Group(children),
	)
}