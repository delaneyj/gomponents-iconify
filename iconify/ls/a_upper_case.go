package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M547 762h-81l-89-242H170L82 762H0L273 0zM273 236l-77 214h155z"/>`),
		g.Group(children),
	)
}