package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 4L6 2H2L1 4H0v11h16V4H7zm8 10H1V5h.6l1-2h2.6l1.2 2H15v9z"/>`),
		g.Group(children),
	)
}