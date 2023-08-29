package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fingerprint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 43V22a6 6 0 0 1 12 0v21"/><path d="M12 40V22c0-6.627 5.373-12 12-12s12 5.373 12 12v18"/><path d="M6 35V22c0-9.941 8.059-18 18-18s18 8.059 18 18v13m-18 9V31m0-6.375v-2.75"/></g>`),
		g.Group(children),
	)
}