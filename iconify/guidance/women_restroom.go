package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomenRestroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10 24v-5.5H6.5v-.25l.072-.15A25 25 0 0 0 9 7.35v-.328a8.58 8.58 0 0 1 3-.523c1.288 0 2.311.266 3 .523v.328c0 3.72.83 7.391 2.428 10.749l.072.15v.25H14V24M11.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}