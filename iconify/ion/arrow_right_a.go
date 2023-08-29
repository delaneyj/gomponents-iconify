package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M448.5 256.5l-192-192v112h-192v160h192v112z" fill="currentColor"/>`),
		g.Group(children),
	)
}