package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLayoutCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 12h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5zM9 5h6v6H9V5zm12.5 6h-3a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1zm-19-3h3a.5.5 0 0 0 0-1h-3a.5.5 0 0 0 0 1zm16 0h3a.5.5 0 0 0 0-1h-3a.5.5 0 0 0 0 1zm3 7h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1zm-19-3h3a.5.5 0 0 0 0-1h-3a.5.5 0 0 0 0 1zm11 7h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}