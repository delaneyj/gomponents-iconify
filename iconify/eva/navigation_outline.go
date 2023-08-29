package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaNavigationOutline0"><g id="evaNavigationOutline1"><path id="evaNavigationOutline2" fill="currentColor" d="M20 20a.94.94 0 0 1-.55-.17L12 14.9l-7.45 4.93a1 1 0 0 1-1.44-1.28l8-16a1 1 0 0 1 1.78 0l8 16a1 1 0 0 1-.23 1.2A1 1 0 0 1 20 20Zm-8-7.3a1 1 0 0 1 .55.17l4.88 3.23L12 5.24L6.57 16.1l4.88-3.23a1 1 0 0 1 .55-.17Z"/></g></g>`),
		g.Group(children),
	)
}