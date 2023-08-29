package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpFeaturedPlayList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3H1v18h22V3zm-11 8H3V9h9v2zm0-4H3V5h9v2z"/>`),
		g.Group(children),
	)
}