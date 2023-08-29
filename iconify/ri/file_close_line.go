package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCloseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22v-2H5V4h10v4h4v5h2V7l-5-5H3.998A.995.995 0 0 0 3 2.992v18.016a1 1 0 0 0 .993.992H12Zm9.536-.879L19.414 19l2.122-2.121l-1.415-1.415l-2.12 2.122l-2.122-2.122l-1.414 1.415l2.12 2.12l-2.12 2.122l1.414 1.414L18 20.415l2.121 2.12l1.415-1.414Z"/>`),
		g.Group(children),
	)
}