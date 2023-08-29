package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStraight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.5 4v6h10.159l-1.197-1.71a2.25 2.25 0 0 1 0-2.58L14.659 4H4.5Zm-1-2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h13.08a1 1 0 0 0 .819-1.573L15.1 7.143a.25.25 0 0 1 0-.286l2.3-3.284A1 1 0 0 0 16.579 2H3.5Z"/><path d="M3.5 3a1 1 0 0 1 1 1v13a1 1 0 1 1-2 0V4a1 1 0 0 1 1-1Z"/></g>`),
		g.Group(children),
	)
}