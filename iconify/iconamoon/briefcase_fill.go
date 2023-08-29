package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 2a3 3 0 0 0-3 3v1H4a1 1 0 0 0-1 1v7a2.99 2.99 0 0 0 1 2.236V19a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-2.764c.614-.55 1-1.348 1-2.236V7a1 1 0 0 0-1-1h-4V5a3 3 0 0 0-3-3h-2Zm3 4V5a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v1h4ZM6 19v-2h12v2a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1Zm6-8.5a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V12a1.5 1.5 0 0 0-1.5-1.5H12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}