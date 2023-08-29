package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStraightOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.5 4v6h10.159l-1.197-1.71a2.25 2.25 0 0 1 0-2.58L14.659 4H4.5Zm-1-2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h13.08a1 1 0 0 0 .819-1.573L15.1 7.143a.25.25 0 0 1 0-.286l2.3-3.284A1 1 0 0 0 16.579 2H3.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3.5 3a1 1 0 0 1 1 1v13a1 1 0 1 1-2 0V4a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}