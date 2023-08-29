package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M3.5 8a1 1 0 0 1 0-2h14.423a1 1 0 1 1 0 2H3.5Zm0 3.75a1 1 0 1 1 0-2h14.423a1 1 0 1 1 0 2H3.5Zm0 3.25a1 1 0 1 1 0-2h14.423a1 1 0 1 1 0 2H3.5Z" opacity=".8"/><path d="M2.5 7a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Zm0 3.25a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Zm0 3.25a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Z"/></g>`),
		g.Group(children),
	)
}