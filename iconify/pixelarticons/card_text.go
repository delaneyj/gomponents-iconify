package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4H2v16h20V4H4zm0 2h16v12H4V6zm2 2h12v2H6V8zm0 4h10v2H6v-2z"/>`),
		g.Group(children),
	)
}