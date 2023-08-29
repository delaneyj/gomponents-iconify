package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaNavigation2Outline0"><g id="evaNavigation2Outline1"><path id="evaNavigation2Outline2" fill="currentColor" d="M13.67 22h-.06a1 1 0 0 1-.92-.8L11 13l-8.2-1.69a1 1 0 0 1-.12-1.93l16-5.33A1 1 0 0 1 20 5.32l-5.33 16a1 1 0 0 1-1 .68Zm-6.8-11.9l5.19 1.06a1 1 0 0 1 .79.78l1.05 5.19l3.52-10.55Z"/></g></g>`),
		g.Group(children),
	)
}