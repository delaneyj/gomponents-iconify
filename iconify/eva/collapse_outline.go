package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCollapseOutline0"><g id="evaCollapseOutline1"><path id="evaCollapseOutline2" fill="currentColor" d="M19 9h-2.58l3.29-3.29a1 1 0 1 0-1.42-1.42L15 7.57V5a1 1 0 0 0-1-1a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2Zm-9 4H5a1 1 0 0 0 0 2h2.57l-3.28 3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L9 16.42V19a1 1 0 0 0 1 1a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1Z"/></g></g>`),
		g.Group(children),
	)
}