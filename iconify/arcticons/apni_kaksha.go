package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApniKaksha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.068 42.5l13.012-37l11.695 36.551M12.926 22.998l20.272-8.914m-8.271 9.692l17.005-7.957"/>`),
		g.Group(children),
	)
}