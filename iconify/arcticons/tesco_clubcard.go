package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TescoClubcard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="27.56" x="4.5" y="10.22" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.5" ry="3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.473 26.897h23.054M18.23 21.103h11.54"/>`),
		g.Group(children),
	)
}