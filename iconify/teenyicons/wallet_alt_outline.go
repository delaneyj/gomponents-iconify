package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M13.5 6V2.5a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V9m.93-3.5H9.5a2 2 0 1 0 0 4h4.93a.07.07 0 0 0 .07-.07V5.57a.07.07 0 0 0-.07-.07Z"/>`),
		g.Group(children),
	)
}