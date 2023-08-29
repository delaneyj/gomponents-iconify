package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.502 12.033l-4.241-2.458l2.138-5.131A1.003 1.003 0 0 0 14.505 3a1.004 1.004 0 0 0-.622.214l-.07.06l-7.5 7.1a1.002 1.002 0 0 0 .185 1.592l4.242 2.46l-2.163 5.19a.999.999 0 0 0 1.611 1.11l7.5-7.102a1.002 1.002 0 0 0-.186-1.591z"/>`),
		g.Group(children),
	)
}