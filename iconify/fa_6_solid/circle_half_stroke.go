package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfStroke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 256c0-106-86-192-192-192v384c106 0 192-86 192-192zM0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0z"/>`),
		g.Group(children),
	)
}