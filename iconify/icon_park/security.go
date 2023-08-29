package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Security(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 26V22C42 12.0589 33.9411 4 24 4V4C14.0589 4 6 12.0589 6 22V26"/><path d="M32 27V12C32 7.58172 28.4183 4 24 4V4C19.5817 4 16 7.58172 16 12V27"/><path d="M24 24V38C24 41.3137 26.6863 44 30 44V44C33.3137 44 36 41.3137 36 38V35.1111"/><path d="M12 24H42"/></g>`),
		g.Group(children),
	)
}