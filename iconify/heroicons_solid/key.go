package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 8a6 6 0 0 1-7.743 5.743L10 14l-1 1l-1 1H6v2H2v-4l4.257-4.257A6 6 0 1 1 18 8Zm-6-4a1 1 0 1 0 0 2a2 2 0 0 1 2 2a1 1 0 1 0 2 0a4 4 0 0 0-4-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}