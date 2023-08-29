package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.281 7.281L11.5 19.063L8.719 16.28L7.28 17.72l2.782 2.781L8 22.563L1.719 16.28L.28 17.72l7 7l.719.687l.719-.687l2.781-2.782l2.781 2.782l.719.687l.719-.687l15.906-16l-1.438-1.438L15 22.563L12.937 20.5L24.72 8.719z"/>`),
		g.Group(children),
	)
}