package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 98h-42V56a6 6 0 0 0-6-6H46V40a6 6 0 0 0-12 0v176a6 6 0 0 0 12 0v-10h90a6 6 0 0 0 6-6v-42h74a6 6 0 0 0 6-6v-48a6 6 0 0 0-6-6Zm-54-36v36H46V62Zm-32 132H46v-36h84Zm80-48H46v-36h164Z"/>`),
		g.Group(children),
	)
}