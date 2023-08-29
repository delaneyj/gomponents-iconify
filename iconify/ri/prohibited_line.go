package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.094 5.68L18.32 16.906A8 8 0 0 0 7.094 5.68Zm9.812 12.64L5.68 7.094A8 8 0 0 0 16.906 18.32ZM4.929 4.929A9.972 9.972 0 0 1 12 2c5.523 0 10 4.477 10 10a9.972 9.972 0 0 1-2.929 7.071A9.972 9.972 0 0 1 12 22C6.477 22 2 17.523 2 12a9.972 9.972 0 0 1 2.929-7.071Z"/>`),
		g.Group(children),
	)
}