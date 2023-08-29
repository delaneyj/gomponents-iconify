package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStraightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.5 7v6h10.159l-1.197-1.71a2.25 2.25 0 0 1 0-2.58L17.659 7H7.5Zm-1-2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h13.08a1 1 0 0 0 .819-1.573L18.1 10.143a.25.25 0 0 1 0-.286l2.3-3.284A1 1 0 0 0 19.579 5H6.5Z"/><path d="M6.5 6a1 1 0 0 1 1 1v13a1 1 0 1 1-2 0V7a1 1 0 0 1 1-1Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}