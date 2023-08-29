package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M15 15L24 29L33 15"/><path stroke="#fff" d="M24 28V37"/></g>`),
		g.Group(children),
	)
}