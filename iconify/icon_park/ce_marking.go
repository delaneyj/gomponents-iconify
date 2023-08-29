package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeMarking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M22 42C12.0589 42 4 33.9411 4 24C4 14.0589 12.0589 6 22 6"/><path d="M44 42C34.0589 42 26 33.9411 26 24C26 14.0589 34.0589 6 44 6"/><path d="M26 24H37"/></g>`),
		g.Group(children),
	)
}