package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polymer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m395 21l96 171l-96 171h-86l96-171l-55-99l-169 270H96L0 192L96 21h85L85 192l56 99L309 21h86z"/>`),
		g.Group(children),
	)
}