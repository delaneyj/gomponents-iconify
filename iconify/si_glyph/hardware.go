package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hardware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.007 4.014a2.977 2.977 0 0 0 0 5.954a2.977 2.977 0 0 0 0-5.954zm0 5.021a2.046 2.046 0 0 1 0-4.093a2.046 2.046 0 0 1 0 4.093zM6 14h3.965v2H6zm0-3h3.964v1.959H6zM2.473 4.53l2.1-3.066l1.211.83l-2.1 3.065zm-2.55-1.626l2.1-3.066L3.23.665l-2.1 3.066zm11.673-1.733h3.734v3.457h-3.734z"/>`),
		g.Group(children),
	)
}