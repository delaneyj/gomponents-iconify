package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 9a.5.5 0 0 0-.5.5V14H3V9.5a.5.5 0 0 0-1 0v5a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}