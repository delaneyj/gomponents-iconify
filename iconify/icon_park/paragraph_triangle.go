package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 9H42"/><path d="M24 19H42"/><path d="M6 29H42"/><path d="M6 39H42"/><path fill="#2F88FF" d="M6 9.76619C6 8.9889 6.84797 8.50878 7.5145 8.9087L14.5708 13.1425C15.2182 13.5309 15.2182 14.4691 14.5708 14.8575L7.5145 19.0913C6.84797 19.4912 6 19.0111 6 18.2338V9.76619Z"/></g>`),
		g.Group(children),
	)
}