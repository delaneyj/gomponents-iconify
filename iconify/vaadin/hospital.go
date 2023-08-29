package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 4V0H8v4H0v12h6v-3h4v3h6V4h-1zM4 11H2V9h2v2zm0-3H2V6h2v2zm3 3H5V9h2v2zm0-3H5V6h2v2zm3-5V2h1V1h1v1h1v1h-1v1h-1V3h-1zm1 8H9V9h2v2zm0-3H9V6h2v2zm3 3h-2V9h2v2zm0-3h-2V6h2v2z"/>`),
		g.Group(children),
	)
}