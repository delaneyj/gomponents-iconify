package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorshipChristian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 30h-2a2.002 2.002 0 0 1-2-2V14H8a2.002 2.002 0 0 1-2-2v-2a2.002 2.002 0 0 1 2-2h5V4a2.002 2.002 0 0 1 2-2h2a2.002 2.002 0 0 1 2 2v4h5a2.002 2.002 0 0 1 2 2v2a2.002 2.002 0 0 1-2 2h-5v14a2.002 2.002 0 0 1-2 2ZM8 10v2h7v16h2V12h7v-2h-7V4h-2v6Z"/>`),
		g.Group(children),
	)
}