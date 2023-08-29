package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UniversalAccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2zm0 3.33A1.67 1.67 0 1 1 10.33 7A1.67 1.67 0 0 1 12 5.33zm3.33 12.5l-1.66.84l-1.39-3.89h-.56l-1.39 3.89l-1.66-.84l1.66-4.72v-1.66L7 10.33l.56-1.66l3.33 1.11h2.22l3.33-1.11l.56 1.66l-3.33 1.12v1.66z"/>`),
		g.Group(children),
	)
}