package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 12h2l-2.5 3l-2.5-3h2V4h-2l2.5-3L16 4h-2zM10 1v4L9 3H6v11h2v1H2v-1h2V3H1L0 5V1z"/>`),
		g.Group(children),
	)
}