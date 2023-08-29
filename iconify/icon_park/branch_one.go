package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 33V15"/><rect width="28" height="6" x="10" y="9" fill="#2F88FF"/><path d="M8 32L14 25H33.9743L40 32"/><rect width="8" height="8" x="4" y="33" fill="#2F88FF"/><rect width="8" height="8" x="20" y="33" fill="#2F88FF"/><rect width="8" height="8" x="36" y="33" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}