package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerCord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 4.414L14.586 3l-2.793 2.793l-1.586-1.586L13 1.414L11.586 0L8.793 2.793L7 1L5.646 2.353l8 8L15 9l-1.793-1.793L16 4.414zm-3.593 6.114L5.472 3.593C3.975 5.388 2.276 8.163 3.45 10.55l-2.066 2.066a1.254 1.254 0 0 0 0 1.768l.232.232a1.254 1.254 0 0 0 1.768 0L5.45 12.55c2.387 1.174 5.161-.524 6.957-2.022z"/>`),
		g.Group(children),
	)
}