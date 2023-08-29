package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.556 3.444a4 4 0 0 1 0 5.656l-8.2 8.2a4 4 0 0 1-2.386 1.148l-3.38.374l-2.297 2.3a1 1 0 0 1-1.414-1.415l2.298-2.299l.375-3.378A4 4 0 0 1 6.7 11.645l8.2-8.2a4 4 0 0 1 5.658 0Zm-4.242 1.414l-8.2 8.2a2 2 0 0 0-.574 1.193l-.276 2.485l2.485-.276a2 2 0 0 0 1.193-.574l.422-.422L9.95 14.05l1.414-1.414l1.414 1.414l1.414-1.414l-1.414-1.414l1.415-1.414l1.414 1.414l1.414-1.415l-1.414-1.414L17.02 6.98l1.414 1.414l.707-.707a2 2 0 0 0-2.828-2.828Z"/>`),
		g.Group(children),
	)
}