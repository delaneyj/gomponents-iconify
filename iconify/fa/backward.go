package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1523 13q19-19 32-13t13 32v1472q0 26-13 32t-32-13L813 813q-9-9-13-19v710q0 26-13 32t-32-13L45 813q-19-19-19-45t19-45L755 13q19-19 32-13t13 32v710q4-10 13-19z"/>`),
		g.Group(children),
	)
}