package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PokerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.494 1.01H4.412C3.082 1.01 2 1.107 2 2.454v13.101C2 16.903 3.082 17 4.412 17h6.082c1.33 0 2.412-.098 2.412-1.445V2.454c0-1.348-1.082-1.444-2.412-1.444zM7.489 13.135L4.748 9.041l2.824-4.115l2.742 4.093l-2.825 4.116z"/>`),
		g.Group(children),
	)
}