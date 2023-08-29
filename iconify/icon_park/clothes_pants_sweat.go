package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPantsSweat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 19L33 38H42L38 4H10L6 38H15L24 19Z"/><path fill="#2F88FF" d="M34 38L35 44H41V38H34Z"/><path fill="#2F88FF" d="M13 44H7V38H14L13 44Z"/><path d="M24 4L28 11.5"/><path d="M24 4L20 11.5"/></g>`),
		g.Group(children),
	)
}