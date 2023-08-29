package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNibStraightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220.54 124.77a1.91 1.91 0 0 0-.15-.28L190 70.42V32a14 14 0 0 0-14-14H80a14 14 0 0 0-14 14v38.44l-30.4 54.05a1.91 1.91 0 0 0-.15.28a14 14 0 0 0 1.27 14.5a.76.76 0 0 1 .08.11l86.44 112.28a6 6 0 0 0 9.51 0l86.43-112.28a.76.76 0 0 1 .08-.11a14 14 0 0 0 1.28-14.5ZM80 30h96a2 2 0 0 1 2 2v34H78V32a2 2 0 0 1 2-2Zm48 116a14 14 0 1 1 14-14a14 14 0 0 1-14 14Zm81.63-13.88L134 230.38v-73.09a26 26 0 1 0-12 0v73.07l-75.63-98.24a2 2 0 0 1-.2-1.93L75.52 78h105l29.34 52.19a2 2 0 0 1-.23 1.93Z"/>`),
		g.Group(children),
	)
}