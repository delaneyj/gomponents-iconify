package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M41.708 23.794L24.058 6.086L6.294 23.792a1 1 0 1 0 1.412 1.416L11 21.925V41a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1v-8a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v8a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V21.904l3.292 3.302a1 1 0 1 0 1.416-1.412Z"/>`),
		g.Group(children),
	)
}