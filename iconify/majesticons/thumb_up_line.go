package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.926 3.66a2.254 2.254 0 0 1 3.817 1.965l-.563 3.378A5 5 0 0 1 21 14v2a5 5 0 0 1-5 5H6a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h1.586l5.34-5.34zM7 11H6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h1v-8zm2 8h7a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3h-3a1 1 0 1 1 0-2h1.153l.617-3.704a.254.254 0 0 0-.43-.222L9 10.414V19z"/></g>`),
		g.Group(children),
	)
}