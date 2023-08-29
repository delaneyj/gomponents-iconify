package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLayoutLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 12h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5zM3 5h6v6H3V5zm9.5 3h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1zm9 7h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1zm-8 4h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1zm8-8h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}