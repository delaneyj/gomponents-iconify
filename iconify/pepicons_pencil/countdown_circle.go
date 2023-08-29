package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CountdownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 6a.5.5 0 0 1 .5-.5A7.5 7.5 0 1 1 6 13a.5.5 0 0 1 1 0a6.5 6.5 0 1 0 6.5-6.5a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M11 6.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-2.5 1.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM7 10.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path fill-rule="evenodd" d="M8.947 14.224a.5.5 0 0 0-.223-.671l-2-1a.5.5 0 0 0-.448.894l2 1a.5.5 0 0 0 .671-.223Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.854 12.646a.5.5 0 0 1 0 .708l-1.5 1.5a.5.5 0 1 1-.708-.707l1.5-1.5a.5.5 0 0 1 .708 0ZM13.5 9.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 13a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}