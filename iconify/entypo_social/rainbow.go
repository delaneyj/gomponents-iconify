package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 5C4.477 5 0 9.477 0 15h1.5a8.5 8.5 0 0 1 17 0H20c0-5.523-4.477-10-10-10zm0 6a4 4 0 0 0-4 4h1.5a2.5 2.5 0 1 1 5 0H14a4 4 0 0 0-4-4zm0-3a7 7 0 0 0-7 7h1.5a5.5 5.5 0 1 1 11 0H17a7 7 0 0 0-7-7z"/>`),
		g.Group(children),
	)
}