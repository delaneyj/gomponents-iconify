package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h20v16H2V4zm2 2v12h16V6H4zm6 4h8v6h-8v-6zm2 2v2h4v-2h-4z"/>`),
		g.Group(children),
	)
}