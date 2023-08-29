package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.4 3a.5.5 0 0 0-.434.25l-2.6 4.5a.5.5 0 0 0 0 .5l2.6 4.5A.5.5 0 0 0 5.4 13h5.2a.5.5 0 0 0 .433-.25l2.6-4.5a.5.5 0 0 0 0-.5l-2.6-4.5A.5.5 0 0 0 10.6 3H5.4Zm-1.3-.25A1.5 1.5 0 0 1 5.4 2h5.2a1.5 1.5 0 0 1 1.298.75l2.6 4.5a1.5 1.5 0 0 1 0 1.5l-2.6 4.5A1.5 1.5 0 0 1 10.6 14H5.4a1.5 1.5 0 0 1-1.298-.75l-2.6-4.5a1.5 1.5 0 0 1 0-1.5l2.6-4.5Z"/>`),
		g.Group(children),
	)
}