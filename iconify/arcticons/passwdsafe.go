package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passwdsafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.167 30.167H5.5L17.833 42.5H5.5v-37l12.333 12.333H5.5L30.167 42.5m0 0H42.5L30.167 30.167"/>`),
		g.Group(children),
	)
}