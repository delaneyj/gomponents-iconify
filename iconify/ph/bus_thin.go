package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 36H72a28 28 0 0 0-28 28v144a12 12 0 0 0 12 12h24a12 12 0 0 0 12-12v-20h72v20a12 12 0 0 0 12 12h24a12 12 0 0 0 12-12V64a28 28 0 0 0-28-28ZM52 180v-64h152v64Zm152-72H52V76h152ZM84 208a4 4 0 0 1-4 4H56a4 4 0 0 1-4-4v-20h32Zm116 4h-24a4 4 0 0 1-4-4v-20h32v20a4 4 0 0 1-4 4Zm4-144H52v-4a20 20 0 0 1 20-20h112a20 20 0 0 1 20 20Zm-104 80a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm72 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm72-68v24a4 4 0 0 1-8 0V80a4 4 0 0 1 8 0ZM20 80v24a4 4 0 0 1-8 0V80a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}