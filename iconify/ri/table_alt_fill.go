package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableAltFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14V3H3a1 1 0 0 0-1 1v10h5Zm8 0V3H9v11h6Zm7 0V4a1 1 0 0 0-1-1h-4v11h5Zm-1 7a1 1 0 0 0 1-1v-4H2v4a1 1 0 0 0 1 1h18Z"/>`),
		g.Group(children),
	)
}