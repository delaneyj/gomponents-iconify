package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGoogleTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.298 25.5a2.987 2.987 0 0 0 3 3a2.895 2.895 0 0 0 2.889-3v-3a2.966 2.966 0 0 0-2.89-3a3.06 3.06 0 0 0-3 3Zm-4.484-4.762l2.25-1.238m0 0v9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h26l9-9v-26a2 2 0 0 0-2-2Zm-7 0v37m-19-28v28m28-9h-37m37-19h-28"/>`),
		g.Group(children),
	)
}