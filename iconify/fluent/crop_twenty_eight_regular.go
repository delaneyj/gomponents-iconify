package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2.75a.75.75 0 0 0-1.5 0V6.5H2.75a.75.75 0 0 0 0 1.5H6.5v9.75a3.75 3.75 0 0 0 3.75 3.75H20v3.75a.75.75 0 0 0 1.5 0V21.5h3.75a.75.75 0 0 0 0-1.5h-15A2.25 2.25 0 0 1 8 17.75v-15Zm12 7.5V19h1.5v-8.75a3.75 3.75 0 0 0-3.75-3.75H9V8h8.75A2.25 2.25 0 0 1 20 10.25Z"/>`),
		g.Group(children),
	)
}