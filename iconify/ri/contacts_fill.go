package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22a8 8 0 1 1 16 0H2Zm8-9c-3.315 0-6-2.685-6-6s2.685-6 6-6s6 2.685 6 6s-2.685 6-6 6Zm10 4h4v2h-4v-2Zm-3-5h7v2h-7v-2Zm2-5h5v2h-5V7Z"/>`),
		g.Group(children),
	)
}