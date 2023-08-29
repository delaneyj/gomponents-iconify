package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4.25h-1.25V3a.75.75 0 0 0-1.5 0v1.25h-4.5V3a.75.75 0 0 0-1.5 0v1.25H7A2.75 2.75 0 0 0 4.25 7v11A2.75 2.75 0 0 0 7 20.75h10A2.75 2.75 0 0 0 19.75 18V7A2.75 2.75 0 0 0 17 4.25ZM7 5.75h1.25V7a.75.75 0 0 0 1.5 0V5.75h4.5V7a.75.75 0 0 0 1.5 0V5.75H17A1.25 1.25 0 0 1 18.25 7v2.75H5.75V7A1.25 1.25 0 0 1 7 5.75Zm10 13.5H7A1.25 1.25 0 0 1 5.75 18v-6.75h12.5V18A1.25 1.25 0 0 1 17 19.25Z"/>`),
		g.Group(children),
	)
}