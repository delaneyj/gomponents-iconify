package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 8h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1zm11 9h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1zm8-5h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}