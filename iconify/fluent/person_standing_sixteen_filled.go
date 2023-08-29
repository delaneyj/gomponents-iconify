package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonStandingSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 4.5a1.75 1.75 0 1 0 .001-3.499A1.75 1.75 0 0 0 7.5 4.5Zm3.5 2v3c0 .551-.448 1-1 1V14a1 1 0 0 1-2 0v-3a.5.5 0 0 0-1 0v3a1 1 0 0 1-2 0v-3.5c-.552 0-1-.449-1-1v-3c0-.916.623-1.682 1.464-1.918A2.735 2.735 0 0 0 7.5 5.5c.81 0 1.532-.359 2.036-.918A1.997 1.997 0 0 1 11 6.5Z"/>`),
		g.Group(children),
	)
}