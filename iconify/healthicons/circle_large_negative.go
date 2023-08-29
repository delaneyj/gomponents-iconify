package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLargeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCircleLargeNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0v48h48V0H0Zm24 4a20 20 0 1 1 0 40a20 20 0 0 1 0-40Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCircleLargeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}