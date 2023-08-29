package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 3h256v128l-85 85l85 85v128H0V301l85-85l-85-85V3zm213 309l-85-85l-85 85v75h170v-75zm-85-107l85-85V45H43v75z"/>`),
		g.Group(children),
	)
}