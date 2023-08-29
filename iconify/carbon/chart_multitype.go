package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMultitype(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 22h2v8h-2zm-4-4h2v12h-2zm-4 8h2v4h-2zM9 16a7 7 0 1 0 7 7a7.008 7.008 0 0 0-7-7zm4.899 6H10v-3.899A5.014 5.014 0 0 1 13.899 22zM9 28a5 5 0 0 1-1-9.899V22a2 2 0 0 0 2 2h3.899A5.008 5.008 0 0 1 9 28zm13.535-16l4-6H30V4h-4.535l-4 6H18V2h-2v12a2 2 0 0 0 2 2h12v-2H18v-2z"/><circle cx="11" cy="7" r="1" fill="currentColor"/><circle cx="9" cy="11" r="1" fill="currentColor"/><circle cx="7" cy="5" r="1" fill="currentColor"/><circle cx="5" cy="9" r="1" fill="currentColor"/><circle cx="3" cy="13" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}