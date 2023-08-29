package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConstrainedSurface(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 23a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm18 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM3 5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm18 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-8 7c0-1.246-1.159-2.5-2.572-2.5h-.857C8.152 9.5 7 10.62 7 12c0 1.19.855 2.185 2 2.438c.188.041.38.062.572.062"/><path d="M11 12c0 1.246 1.159 2.5 2.572 2.5h.857C15.848 14.5 17 13.38 17 12c0-1.19-.855-2.186-2-2.438a2.651 2.651 0 0 0-.572-.062M21 19V5M3 19V5m2-2h14M5 21h14"/></g>`),
		g.Group(children),
	)
}