package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 0 1 12.857-5.249l-.714.7A6.5 6.5 0 1 0 13.98 8H8V7h7v.5a7.5 7.5 0 0 1-15 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}