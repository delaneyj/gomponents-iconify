package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M15 24V11.8756L25.5 17.9378L36 24L25.5 30.0622L15 36.1244V24Z"/>`),
		g.Group(children),
	)
}