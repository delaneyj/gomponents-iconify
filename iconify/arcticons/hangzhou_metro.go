package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HangzhouMetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.621 22.36a5.452 5.452 0 0 1 10.904 0v8.723M14.621 16.908v14.175m10.905-8.723a5.452 5.452 0 0 1 10.904 0v8.723"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.364 45.232A21.497 21.497 0 0 1 14.577 4.677m8.673-2.162a21.497 21.497 0 0 1 13.081 39.094m.109-10.521l-.108 10.522M14.631 16.912l-.054-12.235"/>`),
		g.Group(children),
	)
}