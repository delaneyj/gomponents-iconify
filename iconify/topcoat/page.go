package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Page(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22.5 1.5h-14c-2.55 0-3 .561-3 3v32c0 2.49.55 3 3 3h24c2.5 0 3-.47 3-3v-22h-13v-13zm13 10l-10-10v10h10z"/>`),
		g.Group(children),
	)
}