package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10.667 6.25h2.65a.74.74 0 0 1 .033 0H16a.75.75 0 0 1 0 1.5h-2.09l-2.267 8.5h1.69a.75.75 0 0 1 0 1.5H8a.75.75 0 0 1 0-1.5h2.09l2.267-8.5h-1.69a.75.75 0 0 1 0-1.5Z"/><path fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25ZM2.75 12a9.25 9.25 0 1 1 18.5 0a9.25 9.25 0 0 1-18.5 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}