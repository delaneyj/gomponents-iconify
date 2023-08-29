package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFilterCenterFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 9a3 3 0 0 0-3 3a3 3 0 0 0 3 3a3 3 0 0 0 3-3a3 3 0 0 0-3-3m7 10h-4v2h4a2 2 0 0 0 2-2v-4h-2m0-12h-4v2h4v4h2V5a2 2 0 0 0-2-2M5 5h4V3H5a2 2 0 0 0-2 2v4h2m0 6H3v4a2 2 0 0 0 2 2h4v-2H5v-4Z"/>`),
		g.Group(children),
	)
}