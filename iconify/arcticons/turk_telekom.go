package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurkTelekom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 25.684L17.905 10.906v29.555L43.5 25.684zm-25.595-7.715L4.5 10.229v15.479l13.405-7.739z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.257 25.684l12.637 7.295V18.388l-12.637 7.296zm31.488-14.173l-6.88-3.972v7.944l6.88-3.972z"/>`),
		g.Group(children),
	)
}