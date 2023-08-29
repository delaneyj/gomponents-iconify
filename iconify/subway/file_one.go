package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M330 0v128h128L330 0zm-21.3 0H74v512h384V149.3H308.7V0zm64 320h-64v128h-85.3V320h-64L266 213.3L372.7 320z"/>`),
		g.Group(children),
	)
}