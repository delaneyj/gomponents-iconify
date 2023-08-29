package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0zM0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0zm256-96a96 96 0 1 1 0 192a96 96 0 1 1 0-192z"/>`),
		g.Group(children),
	)
}