package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarginRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 2V1h-1V0H0v16h14v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1V9h-1V8h1V7h-1V6h1V5h-1V4h1V3h-1V2h1zm-2 13H1V1h11v14zm3 0h1v1h-1v-1zm-1-1h1v1h-1v-1zm1-1h1v1h-1v-1zm-1-1h1v1h-1v-1zm1-1h1v1h-1v-1zm-1-1h1v1h-1v-1zm1-1h1v1h-1V9zm-1-1h1v1h-1V8zm1-1h1v1h-1V7zm-1-1h1v1h-1V6zm1-1h1v1h-1V5zm-1-1h1v1h-1V4zm1-1h1v1h-1V3zm-1-1h1v1h-1V2zm1-1h1v1h-1V1zm-1-1h1v1h-1V0z"/>`),
		g.Group(children),
	)
}