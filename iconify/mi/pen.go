package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.707 2.293a1 1 0 0 0-1.414 0L14 4.586l-1.293-1.293a1 1 0 0 0-1.414 0l-6 6a1 1 0 0 0 1.414 1.414L12 5.414l.586.586l-2.293 2.293l-7 7A1 1 0 0 0 3 16v4a1 1 0 0 0 1 1h4a1 1 0 0 0 .707-.293l7-7l6-6a1 1 0 0 0 0-1.414l-4-4zm-3 4.414L17 4.414L19.586 7L15 11.586L12.414 9l2.293-2.293zM5 16.414l6-6L13.586 13l-6 6H5v-2.586z"/>`),
		g.Group(children),
	)
}