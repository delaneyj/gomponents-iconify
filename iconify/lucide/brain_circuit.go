package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrainCircuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4.5a2.5 2.5 0 0 0-4.96-.46a2.5 2.5 0 0 0-1.98 3a2.5 2.5 0 0 0-1.32 4.24a3 3 0 0 0 .34 5.58a2.5 2.5 0 0 0 2.96 3.08a2.5 2.5 0 0 0 4.91.05L12 20V4.5ZM16 8V5c0-1.1.9-2 2-2m-6 10h4"/><path d="M12 18h6a2 2 0 0 1 2 2v1M12 8h8m.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-4 5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/><path d="M20.5 21a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-2-18a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/></g>`),
		g.Group(children),
	)
}