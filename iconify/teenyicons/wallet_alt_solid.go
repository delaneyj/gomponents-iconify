package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 12.5V10h.43a.57.57 0 0 0 .57-.57V5.57a.57.57 0 0 0-.57-.57H14V2.5A1.5 1.5 0 0 0 12.5 1h-11A1.5 1.5 0 0 0 0 2.5v10A1.5 1.5 0 0 0 1.5 14h11a1.5 1.5 0 0 0 1.5-1.5ZM14 6v3H9.5a1.5 1.5 0 1 1 0-3H14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}