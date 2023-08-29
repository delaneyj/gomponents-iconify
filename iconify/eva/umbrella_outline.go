package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaUmbrellaOutline0"><g id="evaUmbrellaOutline1"><path id="evaUmbrellaOutline2" fill="currentColor" d="M12 2A10 10 0 0 0 2 12a1 1 0 0 0 1 1h8v6a3 3 0 0 0 6 0a1 1 0 0 0-2 0a1 1 0 0 1-2 0v-6h8a1 1 0 0 0 1-1A10 10 0 0 0 12 2Zm-7.94 9a8 8 0 0 1 15.88 0Z"/></g></g>`),
		g.Group(children),
	)
}