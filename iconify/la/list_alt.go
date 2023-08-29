package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.281 5.281L7 8.563L5.719 7.28L4.28 8.72l2 2l.719.687l.719-.687l4-4zM15 7v2h13V7zm-4.719 6.281L7 16.562l-1.281-1.28l-1.438 1.437l2 2l.719.687l.719-.687l4-4zM15 15v2h13v-2zm-4.719 6.281L7 24.563L5.719 23.28L4.28 24.72l2 2l.719.687l.719-.687l4-4zM15 23v2h13v-2z"/>`),
		g.Group(children),
	)
}