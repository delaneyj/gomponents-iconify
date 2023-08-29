package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostOrg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.049 15.31H7.89v3.774H.049zm11.76 0h7.836v3.774h-7.836zM.043 7.762h19.604v3.774H.043zM.049.214h11.762v3.774H.049zm15.681 0h3.92v3.774h-3.92z"/>`),
		g.Group(children),
	)
}