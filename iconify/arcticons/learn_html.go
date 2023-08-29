package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LearnHtml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M30.7 23.7H17.3l.5 7.3h12.6l-.6 7.9l-5.8 1.6l-5.7-1.6l-.2-3.3"/><path stroke-miterlimit="10" d="M36.4 7.5v9.9h5m-17.1-.1V7.5l4.9 9.9l4.9-9.9v9.9M15.5 7.5H22m-3.3 9.9V7.5m-12.1 0v9.9m6.6-9.9v9.9m-6.6-5h6.6"/></g>`),
		g.Group(children),
	)
}