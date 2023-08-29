package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Zm-1.333 4.25h2.65a.74.74 0 0 1 .033 0H16a.75.75 0 0 1 0 1.5h-2.09l-2.267 8.5h1.69a.75.75 0 0 1 0 1.5H8a.75.75 0 0 1 0-1.5h2.09l2.267-8.5h-1.69a.75.75 0 0 1 0-1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}