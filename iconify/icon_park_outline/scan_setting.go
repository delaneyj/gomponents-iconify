package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" rx="3"/><path d="m34 21l3 3l-3 3m-20-6l-3 3l3 3m13-13l-3-3l-3 3m6 20l-3 3l-3-3"/></g>`),
		g.Group(children),
	)
}