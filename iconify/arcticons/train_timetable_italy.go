package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainTimetableItaly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.25 22.5h31.5M18 37.5l-4.5 6m16.5-6l4.5 6m-21.75-39h22.5a4.5 4.5 0 0 1 4.5 4.5v24a4.5 4.5 0 0 1-4.5 4.5h-22.5a4.5 4.5 0 0 1-4.5-4.5V9a4.5 4.5 0 0 1 4.5-4.5Z"/><circle cx="15.75" cy="30" r=".75" fill="currentColor"/><circle cx="32.25" cy="30" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}