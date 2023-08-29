package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WidgetThreeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M22.25 6.5a4.75 4.75 0 1 0-9.5 0a4.75 4.75 0 0 0 9.5 0Zm-11 11a4.75 4.75 0 1 0-9.5 0a4.75 4.75 0 0 0 9.5 0Z"/><path d="M1.75 6.5a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0Zm11 11a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}