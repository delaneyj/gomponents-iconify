package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0h16.364ZM9.519 7.617l-3.895 3.87a.682.682 0 0 0 .962.968l3.419-3.398l3.49 3.402a.682.682 0 1 0 .952-.976l-3.971-3.87a.682.682 0 0 0-.957.004Z"/>`),
		g.Group(children),
	)
}