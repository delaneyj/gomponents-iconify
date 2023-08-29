package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ok(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1179 72Q929 294 579 822l-115 179Q320 821 0 501l107-107l286 250q150-150 279-271.5T877.5 185T1009 74t77-59l21-14q4 0 11 2t26 19.5t35 49.5z"/>`),
		g.Group(children),
	)
}