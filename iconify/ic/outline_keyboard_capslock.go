package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineKeyboardCapslock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 8.41L16.59 13L18 11.59l-6-6l-6 6L7.41 13L12 8.41zM6 18h12v-2H6v2z"/>`),
		g.Group(children),
	)
}