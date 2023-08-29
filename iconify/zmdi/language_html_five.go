package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageHtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m179 334l91-25l13-138H120l-4-45h171l4-45H66l13 135h155l-5 58l-50 14l-50-14l-4-37H80l7 72zM0 6h357l-32 365l-146 51l-147-51z"/>`),
		g.Group(children),
	)
}