package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M23.963 42L42 27.415L36.996 6l-6.03 12.994H17.993L11.015 6L6 27.415L23.963 42Z"/>`),
		g.Group(children),
	)
}