package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PicOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="24" r="20" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 37L17 28L33 41"/><circle cx="18" cy="17" r="4" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 33L32 23L42 31"/></g>`),
		g.Group(children),
	)
}