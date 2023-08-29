package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M7.5 17a1 1 0 1 1 2 0a.5.5 0 0 0 1 0a1 1 0 1 1 2 0a2.5 2.5 0 0 1-5 0Z"/><path fill-rule="evenodd" d="M18 14.5a3.46 3.46 0 0 0-1.5-2.851V9a6 6 0 0 0-6-6h-1a6 6 0 0 0-6 6v2.649A3.46 3.46 0 0 0 2 14.5A2.5 2.5 0 0 0 4.5 17h11a2.5 2.5 0 0 0 2.5-2.5Zm-3.5-1.618l.59.265l.053.024c.522.237.857.756.857 1.329a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5c0-.573.335-1.092.857-1.33l.053-.023l.59-.265V9a4 4 0 0 1 4-4h1a4 4 0 0 1 4 4v3.882Z" clip-rule="evenodd"/><path d="M9 1.5a1 1 0 0 1 2 0V4a1 1 0 1 1-2 0V1.5ZM1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}