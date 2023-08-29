package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeInfinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 14v-2h2l-9-9l-9 9h2v7a2 2 0 0 0 2 2h2.5"/><path d="M9 21v-6a2 2 0 0 1 2-2h2a2 2 0 0 1 1.75 1.032m.786 3.554a2.123 2.123 0 0 0-2.929 0a1.951 1.951 0 0 0 0 2.828c.809.781 2.12.781 2.929 0c.809-.781-.805.778 0 0l1.46-1.41l1.46-1.419"/><path d="m15.54 17.582l1.46 1.42l1.46 1.41c.809.78-.805-.779 0 0s2.12.781 2.929 0a1.951 1.951 0 0 0 0-2.828a2.123 2.123 0 0 0-2.929 0"/></g>`),
		g.Group(children),
	)
}