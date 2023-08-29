package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkRoundAngleBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m12.792 15.8l1.43-1.432a6.076 6.076 0 0 0 0-8.59a6.067 6.067 0 0 0-8.583 0L2.778 8.643A6.076 6.076 0 0 0 6.732 19"/><path d="M21.222 15.358A6.076 6.076 0 0 0 17.268 5m1.093 13.221a6.067 6.067 0 0 1-8.583 0a6.076 6.076 0 0 1 0-8.589l1.43-1.431"/></g>`),
		g.Group(children),
	)
}