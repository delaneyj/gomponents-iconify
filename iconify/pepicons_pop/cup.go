package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15.5 4h-12a1 1 0 0 0-1 1c0 4.918 3.061 9 7 9s7-4.082 7-9a1 1 0 0 0-1-1Zm-6 8c-2.455 0-4.596-2.57-4.949-6h9.898c-.353 3.43-2.494 6-4.949 6Z" clip-rule="evenodd"/><path d="M4 14.5h11a1 1 0 1 1 0 2H4a1 1 0 1 1 0-2Zm10.024-3.69l.552-1.923c.257.074.539.113.831.113c1.107 0 1.893-.543 1.893-1c0-.457-.786-1-1.893-1V5C17.495 5 19.3 6.248 19.3 8s-1.805 3-3.893 3c-.477 0-.944-.065-1.383-.19Z"/></g>`),
		g.Group(children),
	)
}