package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartMedical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 11h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2Zm6.16-6A6.29 6.29 0 0 0 12 4.41a6.27 6.27 0 0 0-8.16 9.48l6 6.05a3 3 0 0 0 4.24 0l6-6.05A6.27 6.27 0 0 0 20.16 5Zm-1.41 7.46l-6 6a1 1 0 0 1-1.42 0l-6-6a4.29 4.29 0 0 1 0-6a4.27 4.27 0 0 1 6 0a1 1 0 0 0 1.42 0a4.27 4.27 0 0 1 6 0a4.29 4.29 0 0 1 0 6.02Z"/>`),
		g.Group(children),
	)
}