package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.237.075a.5.5 0 0 1 .487-.022l8 4A.5.5 0 0 1 11.5 5h-8a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .237-.425ZM0 8h15V7H0v1Zm3.5 2a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .724.447l8-4A.5.5 0 0 0 11.5 10h-8Z"/>`),
		g.Group(children),
	)
}