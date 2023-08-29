package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 42h40"/><rect width="12" height="20" x="8" y="22" rx="2"/><rect width="20" height="38" x="20" y="4" rx="2"/><path stroke-linecap="round" d="M28 32.008h4m-20 0h4m12-9h4m-4-9h4"/></g>`),
		g.Group(children),
	)
}