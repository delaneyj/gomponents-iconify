package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudArrowUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 128a83.28 83.28 0 0 1-16.8 50.4a4 4 0 1 1-6.4-4.8A76 76 0 1 0 84 128a4 4 0 0 1-8 0a83.45 83.45 0 0 1 4.57-27.27A52 52 0 1 0 72 204h24a4 4 0 0 1 0 8H72A60 60 0 1 1 83.61 93.13A84 84 0 0 1 244 128Zm-89.17-2.83a4 4 0 0 0-5.66 0l-32 32a4 4 0 0 0 5.66 5.66L148 137.66V208a4 4 0 0 0 8 0v-70.34l25.17 25.17a4 4 0 0 0 5.66-5.66Z"/>`),
		g.Group(children),
	)
}