package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOneBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-1.05 0-1.775-.725Q9.5 19.55 9.5 18.5q0-1.05.725-1.775Q10.95 16 12 16q1.05 0 1.775.725q.725.725.725 1.775q0 1.05-.725 1.775Q13.05 21 12 21Z"/>`),
		g.Group(children),
	)
}