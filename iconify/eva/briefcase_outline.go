package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBriefcaseOutline0"><g id="evaBriefcaseOutline1"><path id="evaBriefcaseOutline2" fill="currentColor" d="M19 7h-3V5.5A2.5 2.5 0 0 0 13.5 3h-3A2.5 2.5 0 0 0 8 5.5V7H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-4 2v10H9V9Zm-5-3.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5V7h-4ZM4 18v-8a1 1 0 0 1 1-1h2v10H5a1 1 0 0 1-1-1Zm16 0a1 1 0 0 1-1 1h-2V9h2a1 1 0 0 1 1 1Z"/></g></g>`),
		g.Group(children),
	)
}