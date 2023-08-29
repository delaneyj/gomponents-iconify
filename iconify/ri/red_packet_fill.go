package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedPacketFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.005 5.94a11.985 11.985 0 0 1-6.806 3.863a2.5 2.5 0 0 0-4.388 0A11.985 11.985 0 0 1 3.005 5.94V3.003a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1V5.94Zm0 2.787v12.276a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1V8.727a13.944 13.944 0 0 0 6.63 3.076a2.501 2.501 0 0 0 4.739 0a13.944 13.944 0 0 0 6.63-3.076Z"/>`),
		g.Group(children),
	)
}