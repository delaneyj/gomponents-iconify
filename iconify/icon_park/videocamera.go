package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videocamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="32" height="26" x="4" y="11" rx="2"/><circle cx="20" cy="24" r="6" fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M36 29L44 33V15L36 19"/></g>`),
		g.Group(children),
	)
}