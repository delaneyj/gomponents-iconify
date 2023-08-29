package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarricadeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.556 19H21v2H3v-2h1.444l.89-4h13.333l.889 4ZM17.333 9l.89 4H5.777l.889-4h10.666Zm-.444-2H7.11l.715-3.217A1 1 0 0 1 8.802 3h6.396a1 1 0 0 1 .976.783L16.889 7Z"/>`),
		g.Group(children),
	)
}