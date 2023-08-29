package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftThoughtBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#97d8ea" transform="translate(.11 5.498)"><ellipse cx="49.46" cy="44.28" rx="4.592" ry="3.739"/><ellipse cx="56.946" cy="50.909" rx="2.893" ry="2.235"/><ellipse cx="46.85" cy="24.31" rx="16.787" ry="12.23"/><ellipse cx="28.832" cy="13.299" rx="23.521" ry="13.994"/><ellipse cx="28.833" cy="12.428" rx="18.02" ry="9.893"/><ellipse cx="30.06" cy="29.25" rx="18.02" ry="9.893"/><ellipse cx="16.99" cy="27.01" rx="17.1" ry="9.423"/></g>`),
		g.Group(children),
	)
}