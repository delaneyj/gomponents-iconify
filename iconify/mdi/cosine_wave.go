package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CosineWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2v2c-1.74 0-3 4.58-4.04 8.27c-1.39 5-2.7 9.73-5.96 9.73s-4.57-4.73-5.96-9.73C5 8.58 3.74 4 2 4V2c3.26 0 4.57 4.73 5.96 9.73C9 15.42 10.26 20 12 20c1.74 0 3-4.58 4.04-8.27C17.43 6.73 18.74 2 22 2Z"/>`),
		g.Group(children),
	)
}