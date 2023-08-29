package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThoughtBalloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#97d8ea" transform="translate(.11 5.498)"><ellipse cx="14.177" cy="44.842" rx="4.602" ry="3.746"/><ellipse cx="6.676" cy="51.49" rx="2.899" ry="2.24"/><ellipse cx="16.794" cy="24.828" rx="16.824" ry="12.258"/><ellipse cx="34.851" cy="13.792" rx="23.574" ry="14.02"/><ellipse cx="34.851" cy="12.92" rx="18.06" ry="9.916"/><ellipse cx="33.618" cy="29.778" rx="18.06" ry="9.915"/><ellipse cx="46.721" cy="27.537" rx="17.14" ry="9.444"/></g>`),
		g.Group(children),
	)
}