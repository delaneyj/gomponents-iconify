package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.75 15a4.25 4.25 0 1 0 0-8.5a4.25 4.25 0 0 0 0 8.5Zm0 2a6.25 6.25 0 1 0 0-12.5a6.25 6.25 0 0 0 0 12.5Z" opacity=".2"/><path d="M10 5a5 5 0 1 0 0 10a5 5 0 0 0 0-10Zm-6 5a6 6 0 1 1 12 0a6 6 0 0 1-12 0Z"/></g>`),
		g.Group(children),
	)
}