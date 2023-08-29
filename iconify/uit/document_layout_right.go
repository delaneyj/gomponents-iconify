package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLayoutRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 12h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1zm0-4h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1zm11 11h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1zm8-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1zm0-11h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5zm-.5 7h-6V5h6v6z"/>`),
		g.Group(children),
	)
}