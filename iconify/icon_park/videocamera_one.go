package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocameraOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="31" height="20" x="4" y="21" rx="2"/><rect width="9" height="8" x="20" y="27" stroke-linecap="round" stroke-linejoin="round"/><circle cx="27" cy="13" r="5" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13" cy="13" r="5" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M35 35L44 39V23L35 27"/></g>`),
		g.Group(children),
	)
}