package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarCloseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3V1H7v2H3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-4V1h-2v2H9Zm-5 7h16v9H4v-9Zm0-5h3v1h2V5h6v1h2V5h3v3H4V5Zm5.879 5.964L12 13.086l2.121-2.122l1.415 1.415l-2.122 2.121l2.121 2.121l-1.414 1.414L12 15.915l-2.121 2.12l-1.415-1.414l2.122-2.12l-2.122-2.122l1.415-1.415Z"/>`),
		g.Group(children),
	)
}