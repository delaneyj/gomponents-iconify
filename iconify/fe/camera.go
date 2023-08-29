package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCamera0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCamera1" fill="currentColor"><path id="feCamera2" d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5ZM17 5v2h2V5h-2Zm-5 12a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0-2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></g>`),
		g.Group(children),
	)
}