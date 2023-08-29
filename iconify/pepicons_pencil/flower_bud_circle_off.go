package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerBudCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M16 13a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path fill-rule="evenodd" d="M18.944 16.371a3.502 3.502 0 0 0 0-6.742A3.5 3.5 0 0 0 13 6.55a3.5 3.5 0 0 0-5.944 3.078a3.502 3.502 0 0 0 0 6.742A3.5 3.5 0 0 0 13 19.45a3.5 3.5 0 0 0 5.944-3.078Zm-1.091-6.523a.5.5 0 0 0 .417.666a2.5 2.5 0 0 1 0 4.972a.5.5 0 0 0-.417.666a2.5 2.5 0 0 1-4.436 2.23a.5.5 0 0 0-.833 0a2.5 2.5 0 0 1-4.436-2.23a.5.5 0 0 0-.418-.666a2.5 2.5 0 0 1 0-4.972a.5.5 0 0 0 .417-.666a2.5 2.5 0 0 1 4.436-2.23a.5.5 0 0 0 .833 0a2.5 2.5 0 0 1 4.436 2.23Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}