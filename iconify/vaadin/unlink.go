package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0h1v4H8V0zm0 12h1v4H8v-4zM7 9H3a1 1 0 0 1 0-2h4V5H3a3 3 0 1 0 0 6h4V9zm6-4H9v2h4a1 1 0 0 1 0 2H9v2h4a3 3 0 1 0 0-6zM4.51 15.44L7 12H5.77l-2.08 2.88l.82.56zm7.98 0L10 12h1.23l2.08 2.88l-.82.56zm0-14.45L10 4h1.23l2.08-2.66l-.82-.35zm-7.98 0L7 4H5.77L3.69 1.34l.82-.35z"/>`),
		g.Group(children),
	)
}