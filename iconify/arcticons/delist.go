package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.284 9.855l8.89 5.176L35.56 4.5M12 14.228l10.797 10.441l8.538-4.64M36 23.509L25.438 33.683l-9.154-5.087M12 33.148L22.709 43.5l8.626-5.176"/>`),
		g.Group(children),
	)
}