package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M38 4H14V44L24 26H32L38 4Z"/><path stroke="#fff" d="M35 15H25"/><path stroke="#000" d="M36.636 9L33.3633 21"/></g>`),
		g.Group(children),
	)
}