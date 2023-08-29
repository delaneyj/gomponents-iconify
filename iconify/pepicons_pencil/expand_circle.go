package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M14.354 12.354a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Zm-6 6a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Z"/><path d="M8 18.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1H8Z"/><path d="M8.5 18a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 1 0v4Zm10-6a.5.5 0 0 1-1 0V8a.5.5 0 0 1 1 0v4Z"/><path d="M14 8.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1h-4Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}