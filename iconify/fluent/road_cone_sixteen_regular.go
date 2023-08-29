package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadConeSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.462 1a1 1 0 0 0-.97.757L3.43 14H1.5a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-1.931L9.509 1.757A1 1 0 0 0 8.538 1H7.461Zm4.076 13H4.462l.625-2.5H9.25a.5.5 0 0 0 0-1H5.337l.5-2H8.5a.5.5 0 0 0 0-1H6.087L7.462 2h1.076l3 12Z"/>`),
		g.Group(children),
	)
}