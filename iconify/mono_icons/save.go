package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5a2 2 0 0 1 2-2h11.586A2 2 0 0 1 18 3.586l2.707 2.707A1 1 0 0 1 21 7v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5zm6 14h6v-6H9v6zm8 0h2V7.414l-2-2V7a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V5H5v14h2v-6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v6zM9 5v2h6V5H9z"/>`),
		g.Group(children),
	)
}