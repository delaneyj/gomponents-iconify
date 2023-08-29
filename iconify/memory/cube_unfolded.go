package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeUnfolded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 3v5h10v7h-5v5h-7v-5H0V8h5V3h7m-2 2H7v3h3V5m-3 5v3h3v-3H7m-2 0H2v3h3v-3m12 0v3h3v-3h-3m-2 5h-3v3h3v-3m-3-5v3h3v-3h-3Z"/>`),
		g.Group(children),
	)
}