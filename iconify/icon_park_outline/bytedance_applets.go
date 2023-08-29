package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BytedanceApplets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M24 4v19L8 36m16-13l16 13m-9-16l9-6m-23 6l-9-6m16 17v12"/>`),
		g.Group(children),
	)
}