package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm0 3a7 7 0 0 0-5.106 11.789l.156.16l1.414-1.414a5 5 0 0 1 4.83-8.366l1.564-1.56A6.976 6.976 0 0 0 12 5Zm6.392 4.143l-1.561 1.562a5.008 5.008 0 0 1-1.295 4.83l1.414 1.415A6.978 6.978 0 0 0 19 12a6.975 6.975 0 0 0-.608-2.857Zm-2.15-2.8l-3.725 3.725A2.003 2.003 0 0 0 10 12a2 2 0 1 0 3.932-.517l3.725-3.726l-1.414-1.414Z"/>`),
		g.Group(children),
	)
}