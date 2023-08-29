package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.25 2a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0V3.81l-4.555 4.556a8.75 8.75 0 1 1-1.06-1.06l4.554-4.556H17a.75.75 0 0 1-.75-.75ZM10 6.75a7.25 7.25 0 1 0 0 14.5a7.25 7.25 0 0 0 0-14.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}