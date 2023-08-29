package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnpinFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.971 17.172l-1.414 1.414l-3.535-3.535l-.073.074l-.707 3.535l-1.415 1.415l-4.242-4.243l-4.95 4.95l-1.414-1.414l4.95-4.95l-4.243-4.243l1.414-1.414l3.536-.707l.073-.074l-3.536-3.536l1.414-1.415L20.97 17.172Zm-2.12-4.95l1.34-1.34l.707.707l1.415-1.414l-8.486-8.485l-1.414 1.414l.707.707l-1.34 1.34l7.07 7.072Z"/>`),
		g.Group(children),
	)
}