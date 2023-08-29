package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeMarking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22 42c-9.941 0-18-8.059-18-18S12.059 6 22 6m22 36c-9.941 0-18-8.059-18-18S34.059 6 44 6M26 24h11"/>`),
		g.Group(children),
	)
}