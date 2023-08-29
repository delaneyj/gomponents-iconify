package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 14L34 4l-3.75 3.75l-3.75 3.75L19 19l-7.5 7.5l-3.75 3.75L4 34l10 10l30-30ZM30.25 7.75l-22.5 22.5M9 29l4 4m1-9l6 6m-1-11l4 4m1-9l6 6M29 9l4 4"/>`),
		g.Group(children),
	)
}