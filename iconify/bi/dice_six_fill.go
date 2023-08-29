package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceSixFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 0a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V3a3 3 0 0 0-3-3H3zm1 5.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zm8 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zm1.5 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0zM12 9.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zM5.5 12a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0zM4 9.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z"/>`),
		g.Group(children),
	)
}