package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.394 2.08a1 1 0 0 0-.788 0L5.212 4.822A2 2 0 0 0 4 6.66v6.86a7 7 0 0 0 3.527 6.077l3.977 2.272a1 1 0 0 0 .992 0l3.977-2.272A7 7 0 0 0 20 13.518V6.66a2 2 0 0 0-1.212-1.838l-6.394-2.74Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}