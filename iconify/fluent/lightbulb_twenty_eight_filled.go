package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2a9 9 0 0 0-5.79 15.89c.402.339.679.736.78 1.158l.35 1.452h9.322l.349-1.452c.101-.422.378-.819.78-1.158A9 9 0 0 0 14 2Zm4.301 20H9.7l.362 1.508A3.25 3.25 0 0 0 13.22 26h1.558a3.25 3.25 0 0 0 3.16-2.492L18.301 22Z"/>`),
		g.Group(children),
	)
}