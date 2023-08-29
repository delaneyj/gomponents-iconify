package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.142 5.929c-1.172-1.172-2.707-1.757-4.243-1.757s-3.071.586-4.243 1.757L6 9.586V6H4v7h7v-2H7.414l3.657-3.657a3.975 3.975 0 0 1 2.828-1.172c1.068 0 2.073.416 2.828 1.172a4.004 4.004 0 0 1 0 5.657l-5.364 5.364l1.414 1.414l5.364-5.364a5.999 5.999 0 0 0 .001-8.485z"/>`),
		g.Group(children),
	)
}