package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 1H3L0 4v10.5a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 .5-.5V4l-3-3zm-3 9v3H6v-3H3l5-4l5 4h-3zM2.414 3l1-1h9.171l1 1H2.414z"/>`),
		g.Group(children),
	)
}