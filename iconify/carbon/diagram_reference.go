package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiagramReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4z"/><path fill="currentColor" d="M15 28v-2a9.013 9.013 0 0 0 8.945-8H16a2.002 2.002 0 0 1-2-2V8.055A9.013 9.013 0 0 0 6 17H4A11.012 11.012 0 0 1 15 6a1 1 0 0 1 1 1v9h9a1 1 0 0 1 1 1a11.012 11.012 0 0 1-11 11Z"/><path fill="currentColor" d="M29.006 14h-9.011A1.996 1.996 0 0 1 18 12V3a1.008 1.008 0 0 1 1.02-1A11.012 11.012 0 0 1 30 12.98a1.004 1.004 0 0 1-.994 1.02ZM20 12h7.945A9.018 9.018 0 0 0 20 4.055Z"/>`),
		g.Group(children),
	)
}